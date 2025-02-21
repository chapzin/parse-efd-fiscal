// Package worker fornece uma implementação de pool de workers para execução
// concorrente de tarefas com suporte a cancelamento e timeout.
package worker

import (
	"context"
	"sync"
	"time"
)

// Valores padrão para o pool de workers
const (
	DefaultMaxWorkers  = 3
	DefaultTaskTimeout = 30 * time.Minute
)

// Task representa uma função que será executada pelo worker pool
// e pode retornar um erro.
type Task func() error

// Pool implementa um pool de workers que executam tarefas concorrentemente.
// O número máximo de workers é fixo e as tarefas são distribuídas entre eles.
type Pool struct {
	maxWorkers int
	tasks      chan Task
	results    chan error
	wg         sync.WaitGroup
	ctx        context.Context
}

// NewPool cria um novo pool de workers com o número máximo especificado.
// O contexto é usado para controlar o cancelamento das tarefas.
func NewPool(ctx context.Context, maxWorkers int) *Pool {
	if maxWorkers <= 0 {
		maxWorkers = DefaultMaxWorkers
	}
	return &Pool{
		maxWorkers: maxWorkers,
		tasks:      make(chan Task),
		results:    make(chan error, maxWorkers),
		ctx:        ctx,
	}
}

// Start inicia os workers do pool. Deve ser chamado antes de submeter tarefas.
func (p *Pool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		go p.worker()
	}
}

// worker é a goroutine que executa as tarefas.
// Continua executando até que o contexto seja cancelado ou o canal de tarefas seja fechado.
func (p *Pool) worker() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			if err := task(); err != nil {
				p.results <- err
			}
			p.wg.Done()
		}
	}
}

// Submit adiciona uma nova tarefa ao pool.
// Se o contexto já estiver cancelado, retorna imediatamente com o erro do contexto.
func (p *Pool) Submit(task Task) {
	select {
	case <-p.ctx.Done():
		p.results <- p.ctx.Err()
		return
	default:
		p.wg.Add(1)
		p.tasks <- task
	}
}

// Wait espera todas as tarefas serem concluídas e retorna os erros encontrados.
// Se o contexto for cancelado antes da conclusão, retorna o erro do contexto.
func (p *Pool) Wait() []error {
	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-p.ctx.Done():
		return []error{p.ctx.Err()}
	case <-done:
		close(p.tasks)
		close(p.results)

		var errors []error
		for err := range p.results {
			errors = append(errors, err)
		}
		return errors
	}
}

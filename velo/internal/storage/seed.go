package storage

import "log"

func Seed(store *Engine) error {
	users, _ := store.ListUsers()
	if len(users) > 0 {
		log.Println("📦 Database already seeded, skipping...")
		return nil
	}

	log.Println("🌱 Seeding database with initial data...")

	usersData := []struct {
		id, name, email, password, avatar string
	}{
		{"1", "João Silva", "joao@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=joao"},
		{"2", "Maria Santos", "maria@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=maria"},
		{"3", "Pedro Oliveira", "pedro@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=pedro"},
		{"4", "Ana Costa", "ana@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=ana"},
		{"5", "Lucas Ferreira", "lucas@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=lucas"},
		{"6", "Julia Mendes", "julia@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=julia"},
		{"7", "Rafael Lima", "rafael@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=rafael"},
		{"8", "Camila Souza", "camila@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=camila"},
		{"9", "Bruno Almeida", "bruno@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=bruno"},
		{"10", "Fernanda Rocha", "fernanda@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=fernanda"},
	}

	for _, u := range usersData {
		user := NewUser(u.id, u.name, u.email, u.password)
		user.Avatar = u.avatar
		if err := store.CreateUser(user); err != nil {
			return err
		}
	}
	log.Printf("  ✅ Created %d users\n", len(usersData))

	postsData := []struct {
		id, userID, title, body string
		tags                    []string
	}{
		{"1", "1", "Introdução ao Velo API", "O Velo é um gateway de API de alta performance construído em Go. Ele resolve os 7 problemas fundamentais de qualquer API.", []string{"go", "api", "gateway"}},
		{"2", "1", "Rate Limiting Explorado", "Rate limiting é essencial para proteger sua API de abusos. O Velo implementa token bucket com precisão.", []string{"security", "rate-limit"}},
		{"3", "2", "Cache Inteligente com BadgerDB", "O BadgerDB permite cache de alta performance com persistência em disco. Veja como implementar.", []string{"database", "cache", "performance"}},
		{"4", "2", "Autenticação JWT no Velo", "Aprenda a configurar autenticação JWT e API keys no Velo Gateway.", []string{"auth", "jwt", "security"}},
		{"5", "3", "Load Balancing Round-Robin", "Distribua suas requisições equilibradamente entre múltiplos backends.", []string{"load-balancer", "architecture"}},
		{"6", "3", "Observabilidade com Prometheus", "Métricas detalhadas de cada endpoint usando Prometheus.", []string{"monitoring", "prometheus", "devops"}},
		{"7", "4", "Versionamento de API", "Gerencie versões da sua API usando headers, paths ou query parameters.", []string{"api", "versioning", "best-practices"}},
		{"8", "4", "Documentação Automática OpenAPI", "Gere documentação OpenAPI 3.1 automaticamente.", []string{"docs", "openapi", "swagger"}},
		{"9", "5", "Deploy com Docker", "Containerize seu Velo Gateway com Docker para deploy em qualquer lugar.", []string{"docker", "deploy", "devops"}},
		{"10", "5", "Deploy no Railway", "Deploy gratuito e fácil no Railway com GitHub Actions.", []string{"railway", "deploy", "ci-cd"}},
		{"11", "6", "Performance em Produção", "Benchmarks e dicas para máximo desempenho em produção.", []string{"performance", "benchmarks"}},
		{"12", "6", "HTTPS e TLS", "Configure HTTPS e TLS para segurança em produção.", []string{"security", "tls", "https"}},
		{"13", "7", "Métricas de Latência", "Como medir e otimizar latência da sua API.", []string{"metrics", "latency", "performance"}},
		{"14", "7", "Rate Limiting Adaptativo", "Rate limiting inteligente baseado em padrões de tráfego.", []string{"ai", "rate-limit", "adaptive"}},
		{"15", "8", "Cache Multi-Camada", "Estratégia de cache L1 (RAM) + L2 (BadgerDB) para máxima performance.", []string{"cache", "architecture", "performance"}},
		{"16", "8", "Circuit Breaker", "Implemente circuit breaker para resiliência.", []string{"resilience", "circuit-breaker"}},
		{"17", "9", "WebSocket Support", "Suporte a WebSocket para updates em tempo real.", []string{"websocket", "real-time"}},
		{"18", "9", "gRPC Integration", "Integre backends gRPC no Velo Gateway.", []string{"grpc", "microservices"}},
		{"19", "10", "Multi-Region Deploy", "Deploy em múltiplas regiões para latência mínima.", []string{"architecture", "multi-region"}},
		{"20", "10", "CI/CD com GitHub Actions", "Pipeline completo de CI/CD para seu gateway.", []string{"ci-cd", "github-actions", "devops"}},
	}

	for _, p := range postsData {
		post := NewPost(p.id, p.userID, p.title, p.body, p.tags)
		if err := store.CreatePost(post); err != nil {
			return err
		}
	}
	log.Printf("  ✅ Created %d posts\n", len(postsData))

	commentsData := []struct {
		id, postID, name, email, body string
	}{
		{"1", "1", "Ana", "ana@email.com", "Excelente introdução! Muito bem explicado."},
		{"2", "1", "Bruno", "bruno@email.com", "O Velo parece muito promissor!"},
		{"3", "2", "Carlos", "carlos@email.com", "Rate limiting é crucial para APIs públicas."},
		{"4", "2", "Diana", "diana@email.com", "Como configurar por endpoint?"},
		{"5", "3", "Eduardo", "eduardo@email.com", "BadgerDB é uma ótima escolha!"},
		{"6", "3", "Fernanda", "fernanda@email.com", "Já usei em produção, funciona perfeitamente."},
		{"7", "4", "Gabriel", "gabriel@email.com", "JWT é o padrão nowadays."},
		{"8", "4", "Helena", "helena@email.com", "E para APIs internas? OAuth2 também suporta?"},
		{"9", "5", "Igor", "igor@email.com", "Load balancing salva vidas!"},
		{"10", "5", "Julia", "julia@email.com", "Funciona com Docker Swarm?"},
		{"11", "6", "Kevin", "kevin@email.com", "Prometheus + Grafana = combinação perfeita."},
		{"12", "6", "Larissa", "larissa@email.com", "E métricas customizadas?"},
		{"13", "7", "Marcos", "marcos@email.com", "Versionamento salva de breaking changes."},
		{"14", "7", "Natalia", "natalia@email.com", "Prefiro versionamento por path."},
		{"15", "8", "Oscar", "oscar@email.com", "OpenAPI é essencial para integração."},
		{"16", "8", "Patricia", "patricia@email.com", "Gero automaticamente com Swagger UI."},
		{"17", "9", "Ricardo", "ricardo@email.com", "Docker facilita muito o deploy."},
		{"18", "9", "Sofia", "sofia@email.com", "Multi-stage build é arte!"},
		{"19", "10", "Thiago", "thiago@email.com", "Railway é gratuito para projetos pequenos."},
		{"20", "10", "Ursula", "ursula@email.com", "GitHub Actions é muito poderoso."},
	}

	for _, c := range commentsData {
		comment := NewComment(c.id, c.postID, c.name, c.email, c.body)
		if err := store.CreateComment(comment); err != nil {
			return err
		}
	}
	log.Printf("  ✅ Created %d comments\n", len(commentsData))

	log.Println("🎉 Database seeding complete!")
	return nil
}

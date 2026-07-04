package storage

import (
	"log"

	"github.com/velo-api/velo/pkg/utils"
)

func Seed(store *Engine) error {
	users, _ := store.ListUsers()
	if len(users) > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	log.Println("Seeding database with initial data...")

	usersData := []struct {
		name, email, password, avatar string
	}{
		{"João Silva", "joao@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=joao"},
		{"Maria Santos", "maria@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=maria"},
		{"Pedro Oliveira", "pedro@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=pedro"},
		{"Ana Costa", "ana@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=ana"},
		{"Lucas Ferreira", "lucas@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=lucas"},
		{"Julia Mendes", "julia@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=julia"},
		{"Rafael Lima", "rafael@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=rafael"},
		{"Camila Souza", "camila@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=camila"},
		{"Bruno Almeida", "bruno@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=bruno"},
		{"Fernanda Rocha", "fernanda@email.com", "123456", "https://api.dicebear.com/7.x/avataaars/svg?seed=fernanda"},
	}

	var userIDs []string
	for _, u := range usersData {
		id := utils.GeneratePrefixedID("usr")
		hashedPassword, err := HashPassword(u.password)
		if err != nil {
			log.Printf("Warning: failed to hash password for %s: %v", u.name, err)
			hashedPassword = u.password
		}
		user := NewUser(id, u.name, u.email, hashedPassword)
		user.Avatar = u.avatar
		if err := store.CreateUser(user); err != nil {
			return err
		}
		userIDs = append(userIDs, id)
	}
	log.Printf("  Created %d users\n", len(usersData))

	postsData := []struct {
		title, body string
		tags        []string
		userIdx     int
	}{
		{"Introdução ao Velo API", "O Velo é um gateway de API de alta performance construído em Go. Ele resolve os 7 problemas fundamentais de qualquer API.", []string{"go", "api", "gateway"}, 0},
		{"Rate Limiting Explorado", "Rate limiting é essencial para proteger sua API de abusos. O Velo implementa token bucket com precisão.", []string{"security", "rate-limit"}, 0},
		{"Cache Inteligente com BadgerDB", "O BadgerDB permite cache de alta performance com persistência em disco. Veja como implementar.", []string{"database", "cache", "performance"}, 1},
		{"Autenticação JWT no Velo", "Aprenda a configurar autenticação JWT e API keys no Velo Gateway.", []string{"auth", "jwt", "security"}, 1},
		{"Load Balancing Round-Robin", "Distribua suas requisições equilibradamente entre múltiplos backends.", []string{"load-balancer", "architecture"}, 2},
		{"Observabilidade com Prometheus", "Métricas detalhadas de cada endpoint usando Prometheus.", []string{"monitoring", "prometheus", "devops"}, 2},
		{"Versionamento de API", "Gerencie versões da sua API usando headers, paths ou query parameters.", []string{"api", "versioning", "best-practices"}, 3},
		{"Documentação Automática OpenAPI", "Gere documentação OpenAPI 3.1 automaticamente.", []string{"docs", "openapi", "swagger"}, 3},
		{"Deploy com Docker", "Containerize seu Velo Gateway com Docker para deploy em qualquer lugar.", []string{"docker", "deploy", "devops"}, 4},
		{"Deploy no Railway", "Deploy gratuito e fácil no Railway com GitHub Actions.", []string{"railway", "deploy", "ci-cd"}, 4},
		{"Performance em Produção", "Benchmarks e dicas para máximo desempenho em produção.", []string{"performance", "benchmarks"}, 5},
		{"HTTPS e TLS", "Configure HTTPS e TLS para segurança em produção.", []string{"security", "tls", "https"}, 5},
		{"Métricas de Latência", "Como medir e otimizar latência da sua API.", []string{"metrics", "latency", "performance"}, 6},
		{"Rate Limiting Adaptativo", "Rate limiting inteligente baseado em padrões de tráfego.", []string{"ai", "rate-limit", "adaptive"}, 6},
		{"Cache Multi-Camada", "Estratégia de cache L1 (RAM) + L2 (BadgerDB) para máxima performance.", []string{"cache", "architecture", "performance"}, 7},
		{"Circuit Breaker", "Implemente circuit breaker para resiliência.", []string{"resilience", "circuit-breaker"}, 7},
		{"WebSocket Support", "Suporte a WebSocket para updates em tempo real.", []string{"websocket", "real-time"}, 8},
		{"gRPC Integration", "Integre backends gRPC no Velo Gateway.", []string{"grpc", "microservices"}, 8},
		{"Multi-Region Deploy", "Deploy em múltiplas regiões para latência mínima.", []string{"architecture", "multi-region"}, 9},
		{"CI/CD com GitHub Actions", "Pipeline completo de CI/CD para seu gateway.", []string{"ci-cd", "github-actions", "devops"}, 9},
	}

	var postIDs []string
	for _, p := range postsData {
		id := utils.GeneratePrefixedID("pst")
		post := NewPost(id, userIDs[p.userIdx], p.title, p.body, p.tags)
		if err := store.CreatePost(post); err != nil {
			return err
		}
		postIDs = append(postIDs, id)
	}
	log.Printf("  Created %d posts\n", len(postsData))

	commentsData := []struct {
		postIdx int
		name    string
		email   string
		body    string
	}{
		{0, "Ana", "ana@email.com", "Excelente introdução! Muito bem explicado."},
		{0, "Bruno", "bruno@email.com", "O Velo parece muito promissor!"},
		{1, "Carlos", "carlos@email.com", "Rate limiting é crucial para APIs públicas."},
		{1, "Diana", "diana@email.com", "Como configurar por endpoint?"},
		{2, "Eduardo", "eduardo@email.com", "BadgerDB é uma ótima escolha!"},
		{2, "Fernanda", "fernanda@email.com", "Já usei em produção, funciona perfeitamente."},
		{3, "Gabriel", "gabriel@email.com", "JWT é o padrão nowadays."},
		{3, "Helena", "helena@email.com", "E para APIs internas? OAuth2 também suporta?"},
		{4, "Igor", "igor@email.com", "Load balancing salva vidas!"},
		{4, "Julia", "julia@email.com", "Funciona com Docker Swarm?"},
		{5, "Kevin", "kevin@email.com", "Prometheus + Grafana = combinação perfeita."},
		{5, "Larissa", "larissa@email.com", "E métricas customizadas?"},
		{6, "Marcos", "marcos@email.com", "Versionamento salva de breaking changes."},
		{6, "Natalia", "natalia@email.com", "Prefiro versionamento por path."},
		{7, "Oscar", "oscar@email.com", "OpenAPI é essencial para integração."},
		{7, "Patricia", "patricia@email.com", "Gero automaticamente com Swagger UI."},
		{8, "Ricardo", "ricardo@email.com", "Docker facilita muito o deploy."},
		{8, "Sofia", "sofia@email.com", "Multi-stage build é arte!"},
		{9, "Thiago", "thiago@email.com", "Railway é gratuito para projetos pequenos."},
		{9, "Ursula", "ursula@email.com", "GitHub Actions é muito poderoso."},
	}

	for _, c := range commentsData {
		id := utils.GeneratePrefixedID("cmt")
		comment := NewComment(id, postIDs[c.postIdx], c.name, c.email, c.body)
		if err := store.CreateComment(comment); err != nil {
			return err
		}
	}
	log.Printf("  Created %d comments\n", len(commentsData))

	log.Println("Database seeding complete!")
	return nil
}

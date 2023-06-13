run:
	go run cmd/server/main.go

migrate:
	migrate -path database/postgres/migrations -database "postgres://root:root@host.docker.internal:5432/${table}?sslmode=disable" -verbose ${verbose}

table:
	migrate create -ext sql -dir database/postgres/migrations -seq ${table}

gql:
	cd internal/http/controller && go run github.com/99designs/gqlgen@v0.17.31 generate

test:
	cd tests/integration && go test -v

test-fn:
	cd tests/integration && go test -v -run ${fn}

dataloaden:
	cd internal/http/controller/model && \
	go run github.com/vektah/dataloaden ProjectSkillsLoader int64 []*github.com/arvians-id/go-portfolio/internal/http/controller/model.Skill && \
	go run github.com/vektah/dataloaden WorkExperienceSkillsLoader int64 []*github.com/arvians-id/go-portfolio/internal/http/controller/model.Skill && \
	go run github.com/vektah/dataloaden CategorySkillsLoader int64 []*github.com/arvians-id/go-portfolio/internal/http/controller/model.Skill && \
	go run github.com/vektah/dataloaden SkillsCategoryLoader int64 *github.com/arvians-id/go-portfolio/internal/http/controller/model.CategorySkill && \
	go run github.com/vektah/dataloaden ProjectImagesLoader int64 []*github.com/arvians-id/go-portfolio/internal/http/controller/model.ProjectImage

assets-image:
	cd assets/images && mkdir certificate contact project skill user
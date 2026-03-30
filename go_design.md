<!-- 
┌─────────────────────────────────────────────────┐
│              GO DESIGN PHILOSOPHY               │
├─────────────────────────────────────────────────┤
│                                                 │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐   │
│   │ SIMPLICITY│  │  SPEED    │  │CONCURRENCY│   │
│   │           │  │           │  │           │   │
│   │ No classes│  │ Compiles  │  │ Goroutines│   │
│   │ No inherit│  │ in seconds│  │ Channels  │   │
│   │ No generics│ │ Runs fast │  │ Built-in  │   │
│   │ (until1.18)│ │ Like C    │  │Lightweight│   │
│   └───────────┘  └───────────┘  └───────────┘   │
│                                                 │
│   ┌───────────┐  ┌───────────┐  ┌───────────┐   │
│   │  SAFETY   │  │  TOOLING  │  │ OPINIONS  │   │
│   │           │  │           │  │           │   │
│   │ Type safe │  │ go fmt    │  │ One way to│   │
│   │ GC managed│  │ go test   │  │ do things │   │
│   │ No pointer│  │ go vet    │  │ gofmt     │   │
│   │ arithmetic│  │ go doc    │  │ enforced  │   │
│   └───────────┘  └───────────┘  └───────────┘   │
└─────────────────────────────────────────────────┘

 -->
package builder

type Rag struct {
	LlmName       string
	Frameworkname string
	AgentName     string
	Database      string
}

type OpenIa struct {
	llm       string
	framework string
	vectorDB  string
	agent     string
}

func newOpenIABuilder() *OpenIa {
	return &OpenIa{}
}

func (o *OpenIa) setLlm() {
	o.agent = "chat gpt 3"
}
func (o *OpenIa) setFramework() {
	o.framework = "langchain"
}
func (o *OpenIa) setAgent() {
	o.vectorDB = "pine db"
}
func (o *OpenIa) setVectorDB() {
	o.agent = "doing some task"
}

func (o *OpenIa) getRag() *Rag {
	return &Rag{
		LlmName:       o.llm,
		Frameworkname: o.framework,
		AgentName:     o.agent,
		Database:      o.vectorDB,
	}
}

type Ollama struct {
	llm       string
	framework string
	vectorDB  string
	agent     string
}

func newOllamaLocalBuilder() *Ollama {
	return &Ollama{}
}

func (o *Ollama) setLlm() {
	o.agent = "grok"
}
func (o *Ollama) setFramework() {
	o.framework = "langchain"
}
func (o *Ollama) setAgent() {
	o.vectorDB = "pine db"
}
func (o *Ollama) setVectorDB() {
	o.agent = "doing some task"
}

func (o *Ollama) getRag() *Rag {
	return &Rag{
		LlmName:       o.llm,
		Frameworkname: o.framework,
		AgentName:     o.agent,
		Database:      o.vectorDB,
	}
}

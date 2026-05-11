package builder

type IBuilder interface {
	setLlm()
	setFramework()
	setAgent()
	setVectorDB()
	getRag() *Rag
}

func getBuilder(builderType string) IBuilder {
	if builderType == "openIa" {
		return newOpenIABuilder()
	}

	if builderType == "Ollama" {
		return newOllamaLocalBuilder()
	}
	return nil
}

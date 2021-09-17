generate-protobufs:
	buf lint
	buf breaking --against ".git#branch=main"
	buf generate
unsafe-generate-protobufs:
	buf lint
	buf generate
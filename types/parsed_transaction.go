package types

type ParsedTransaction struct {
	Message    ParsedTransactionMessage `json:"message"`
	Signatures []string                 `json:"signatures"`
}

type ParsedTransactionMessage struct {
	AccountKeys     []ParsedTransactionAccountKey  `json:"accountKeys"`
	Instructions    []ParsedTransactionInstruction `json:"instructions"`
	RecentBlockhash string                         `json:"recentBlockhash"`
}

type ParsedTransactionAccountKey struct {
	Pubkey   string `json:"pubkey"`
	Signer   bool   `json:"signer"`
	Source   string `json:"source"`
	Writable bool   `json:"writable"`
}

type ParsedTransactionInstruction struct {
	Accounts    []any                   `json:"accounts,omitempty"`
	Data        string                  `json:"data,omitempty"`
	ProgramId   string                  `json:"programId"`
	StackHeight any                     `json:"stackHeight"`
	Parsed      ParsedTransactionParsed `json:"parsed,omitempty"`
	Program     string                  `json:"program,omitempty"`
}

type ParsedTransactionParsed struct {
	Info ParsedTransactionParsedInfo `json:"info"`
	Type string                      `json:"type"`
}

type ParsedTransactionParsedInfo struct {
	Destination string `json:"destination"`
	Lamports    int    `json:"lamports"`
	Source      string `json:"source"`
}

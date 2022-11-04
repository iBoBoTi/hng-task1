package domain

type OperationType int64

const (
	Undefined OperationType = iota
	Addition
	Subtraction
	Multiplication
)

func (s OperationType) String() string {
	switch s {
	case Addition:
		return "addition"
	case Subtraction:
		return "subtraction"
	case Multiplication:
		return "multiplication"
	}
	return "unknown defined operation"
}

type Math struct {
	OperationType string `json:"operation_type"`
	FirstOperand  int    `json:"x"`
	SecondOperand int    `json:"y"`
}

type MathOperationResponse struct {
	SlackUserName string `json:"slackUsername"`
	Result        int    `json:"result"`
	OperationType string `json:"operation_type"`
}

func SolveMath(m *Math) *MathOperationResponse {
	var response MathOperationResponse
	response.SlackUserName = "BoBoTi"
	switch m.OperationType {
	case "addition":
		response.Result = m.FirstOperand + m.SecondOperand
		response.OperationType = m.OperationType

		return &response
	case "subtraction":
		response.Result = m.FirstOperand - m.SecondOperand
		response.OperationType = m.OperationType

		return &response
	case "multiplication":
		response.Result = m.FirstOperand + m.SecondOperand
		response.OperationType = m.OperationType

		return &response
	}
	return nil
}

func SetOperationType(opertionType string) string {
	return ""
}

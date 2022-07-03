package openfeature

type Client struct {
	Name string
}

// GetClient returns a new Client.  Name is a unique identifier for this client
func GetClient(name string) Client {
	return Client{Name: name}
}

// AddHooks adds one or more hooks to the client
func (c Client) AddHooks(hooks ...Hook) {

}

// Type represents the type of a flg
type Type int64

const (
	Boolean Type = iota
	String
	Number
	Object
)

type EvaluationDetails struct {
	FlagKey  string
	FlagType Type
	ResolutionDetail
}

// GetBooleanValue return boolean evaluation for flag
func (c Client) GetBooleanValue(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (bool, error) {
	resolution := api.provider.GetBooleanEvaluation(flag, defaultValue, evalCtx, options...)
	return resolution.Value, resolution.Error()
}

// GetStringValue return string evaluation for flag
func (c Client) GetStringValue(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (string, error) {
	resolution := api.provider.GetStringEvaluation(flag, defaultValue, evalCtx, options...)
	return resolution.Value, resolution.Error()
}

// GetNumberValue return number evaluation for flag
func (c Client) GetNumberValue(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (float64, error) {
	resolution := api.provider.GetNumberEvaluation(flag, defaultValue, evalCtx, options...)
	return resolution.Value, resolution.Error()
}

// GetObjectValue return object evaluation for flag
func (c Client) GetObjectValue(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (interface{}, error) {
	resolution := api.provider.GetObjectEvaluation(flag, defaultValue, evalCtx, options...)
	return resolution.Value, resolution.Error()
}

// GetBooleanValueDetails return boolean evaluation for flag
func (c Client) GetBooleanValueDetails(flag string, defaultValue bool, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetBooleanEvaluation(flag, defaultValue, evalCtx, options...)
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Boolean,
		ResolutionDetail: resolution.ResolutionDetail,
	}, resolution.Error()

}

// GetStringValueDetails return string evaluation for flag
func (c Client) GetStringValueDetails(flag string, defaultValue string, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetStringEvaluation(flag, defaultValue, evalCtx, options...)
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         String,
		ResolutionDetail: resolution.ResolutionDetail,
	}, resolution.Error()
}

// GetNumberValueDetails return number evaluation for flag
func (c Client) GetNumberValueDetails(flag string, defaultValue float64, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetNumberEvaluation(flag, defaultValue, evalCtx, options...)
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Number,
		ResolutionDetail: resolution.ResolutionDetail,
	}, resolution.Error()
}

// GetObjectValueDetails return object evaluation for flag
func (c Client) GetObjectValueDetails(flag string, defaultValue interface{}, evalCtx EvaluationContext, options ...EvaluationOption) (EvaluationDetails, error) {
	resolution := api.provider.GetObjectEvaluation(flag, defaultValue, evalCtx, options...)
	return EvaluationDetails{
		FlagKey:          flag,
		FlagType:         Object,
		ResolutionDetail: resolution,
	}, resolution.Error()
}

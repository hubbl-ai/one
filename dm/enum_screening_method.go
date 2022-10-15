package dm

type ScreeningMethod string

const (
	ScreeningMethodPhysicalInspection       ScreeningMethod = "PHS" // Physical Inspection and/or hand search
	ScreeningMethodVisualCheck              ScreeningMethod = "VCK" // Visual check
	ScreeningMethodXRay                     ScreeningMethod = "XRY" // X-ray equipment
	ScreeningMethodExplosiveDetectionSystem ScreeningMethod = "EDS" // Explosive detection system
	ScreeningMethodExplosiveDetectionDogs   ScreeningMethod = "EDS" // Explosive detection dogs
	ScreeningMethodExplosiveTraceDetection  ScreeningMethod = "ETD" // Explosive trace detection equipment - particles or vapor
	ScreeningMethodMetal                    ScreeningMethod = "CMD" // Cargo metal detection
	ScreeningMethodOther                    ScreeningMethod = "AOM" // Subjected to any other means: this entry should be followed by free text specifying what other mean was used to secure the cargo
)

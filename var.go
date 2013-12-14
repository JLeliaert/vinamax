package vinamax

var (
	B_ext         Vector             // External applied field in T
	Dt            float64            // Timestep in s
	Lijst         Particles          // List containing all the Particles
	T             float64            // Time in s
	Alpha         float64            // Gilbert damping constant
	Temp          float64            // Temperature in K
	Ku1           float64            // Uniaxial anisotropy constant in J/m**3
	Msat          float64   = 860000 // Saturation magnetisation in A/m
	Thresholdbeta float64   =0.7            //The threshold value for the FMM
	Universe      node               // The entire universe of the simulation
	FMM	      bool 	=true	 // Calculate demag with FMM method
)

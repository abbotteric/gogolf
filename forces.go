package main

/*
types
*/
type Vector struct {
	x float64
	y float64
}

type Ball struct {
	pos Vector
	vel Vector
}

/*
constants
*/
const M_gb = .04593     //kg
const R_gb = .04267 / 2 //m
const rho = 1.225       //density of air at sea level
const C_d = 0.5         //rough average drag coefficient of a golf ball

/*
initial conditions
*/
var ball = Ball{
	pos: Vector{0, 0},
	vel: Vector{0, 0},
}
var impact = Vector{22.464, 9.076}

func step(b Ball, dt float64, impact_force Vector) Ball {
	var v_new Vector
	var f_calc = impact_force

	// gravity
	f_calc.y += M_gb * G

	// velocity calculation
	v_new.x = b.vel.x + (f_calc.x/M_gb)*dt
	v_new.y = b.vel.y + (f_calc.y/M_gb)*dt

	var p_new = Vector{b.pos.x + v_new.x*dt, b.pos.y + v_new.y*dt}
	return Ball{
		pos: p_new,
		vel: v_new,
	}
}

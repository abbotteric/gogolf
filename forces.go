package main

import (
	"math"
)

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
	var m_v = math.Sqrt(math.Pow(b.vel.x, 2) * math.Pow(b.vel.y, 2)) // magnitude of the velocity

	// gravity
	f_calc.y += M_gb * G

	// drag
	if m_v != 0 {
		// magnitude of the drag
		var m_drag = 0.5 * rho * m_v * C_d * (math.Pi * R_gb * R_gb)
		var f_drag_x = m_drag * math.Sin(math.Atan(b.vel.x/b.vel.y))
		var f_drag_y = m_drag * math.Cos(math.Atan(b.vel.x/b.vel.y))

		// make sure the drag force and the velocity are in opposite directions
		if f_drag_x*b.vel.x > 0 {
			f_drag_x = -1 * f_drag_x
		}
		if f_drag_y*b.vel.y > 0 {
			f_drag_y = -1 * f_drag_y
		}

		// add the drag forces to the ball flight
		f_calc.x += f_drag_x
		f_calc.y += f_drag_y
	}

	// velocity calculation
	v_new.x = b.vel.x + (f_calc.x/M_gb)*dt
	v_new.y = b.vel.y + (f_calc.y/M_gb)*dt

	var p_new = Vector{b.pos.x + v_new.x*dt, b.pos.y + v_new.y*dt}
	return Ball{
		pos: p_new,
		vel: v_new,
	}
}

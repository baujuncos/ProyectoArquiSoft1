import React from 'react'
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';

export const ComponenteColumna = () => {
    return (
        <div className="componente-columna">
            <p><Link className="home-texto" to="/">Home</Link></p>
            <p><Link className="cursos-texto" to="/cursos">Cursos</Link></p>
            <p><Link className="inscripcion-texto" to="/inscripcion">Inscripción</Link></p>
            <p><Link className="miscursos-texto" to="/miscursos">Mis Cursos</Link></p>
        </div>
    )
}
import './App.css'
import { Navigate, Route, Routes } from 'react-router-dom'
import HomePage from './pages/HomePage'
import Startup from './pages/Startup'
import Tournament from './pages/Tournament'
import TournamentDetails from './pages/TournamentDetails'
import Battle from './pages/Battle'

function App() {
  return (
    <Routes>
      <Route path='/' element={<HomePage />} />
      <Route path='/startups' element={<Startup />} />
      <Route path='/torneios' element={<Tournament />} />
      <Route path='/torneios/:id' element={<TournamentDetails />} />
      <Route path='/torneios/battle/:id' element={<Battle />} />

      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}

export default App

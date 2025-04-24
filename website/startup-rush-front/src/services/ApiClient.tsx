import axios, { AxiosInstance, AxiosResponse } from "axios"

export interface ApiResponse<T> {
    message: string
    status:  string
    data:    T
}

export interface HealthApiResponse {
    message: string
    status:  string
}
  
export interface Participant {
    id:           number
    startupId:    number
    tournamentId: number
    score:        number
}

export interface BattleEvent {
    id:     number
    name:   string
    score:  number
  }
  
  export interface BattleWithEvents {
    battle: Battle
    events: BattleEvent[]
  }
  
export interface Battle {
    id:                 number
    tournamentId:       number
    startup1Id:         number
    startup2Id:         number
    score1:             number
    score2:             number
    finished:           boolean
    winnerId?:           number
    phase:              string
    battleChildren1Id?: number
    battleChildren2Id?: number
}
  

export interface TournamentList {
    id:           number
    finished:     boolean
    championId?:  number         
    participants: Participant[]
    battles:      Battle[]
}

export interface Startup {
    id:         number
    name:       string
    slogan:     string
    foundation: Date
}

interface BattleEventSubmission {
    startupId: number
    eventIds: number[]
}

interface SubmitBattleEventsRequest {
    battle: BattleEventSubmission[]
}

class ApiClient {
    private readonly axios: AxiosInstance
  
    constructor() {
      this.axios = axios.create({
        baseURL: import.meta.env.VITE_API_URL,
        headers: { 'Content-Type': 'application/json' },
        timeout: 5000,
      })
      
      this.axios.interceptors.response.use(
        res => res,
        err => {
          return Promise.reject(err)
        }
      )
    }

    async healthCheck(): Promise<string> {
        try {
            const response: AxiosResponse<HealthApiResponse> = await this.axios.get('/api/health')
            return response.data.message
        } catch (error) {
            console.error('API Health Check Error:', error)
            throw error
        }
    }

    async tournaments(): Promise<TournamentList[]> {
        const resp: ApiResponse<TournamentList[]> = await this.axios.get('/api/tournaments')
        return resp.data
    }

    async createStartup(data: any): Promise<any> {
        const resp: AxiosResponse<any> = await this.axios.post('/api/startups', data)
        return resp.status
    }

    async listStartups(): Promise<any> {
        const resp: AxiosResponse<any> = await this.axios.get('/api/startups')
        return resp.data
    }

    async createTournament(data: any): Promise<any> {
        const resp: AxiosResponse<any> = await this.axios.post('/api/tournaments', data)
        return resp.status
    }

    async getBattlesByTournament(id: string): Promise<any> {
        const resp: AxiosResponse<TournamentList> = await this.axios.get(`/api/tournaments/${id}`)
        return resp.data
    }

    async getBattleById(id: string): Promise<any> {
        const resp: AxiosResponse<BattleWithEvents> = await this.axios.get(`/api/tournaments/battle/${id}`)
        return resp.data
    }

    async startTournamentBattles(id: string): Promise<any> {
        const resp: AxiosResponse<any> = await this.axios.get(`/api/tournaments/start/${id}`)
        return resp.data
    }

    async submitBattleEvents(battleId: string, data: SubmitBattleEventsRequest): Promise<any> {
        const resp: AxiosResponse<any> = await this.axios.post(`/api/tournaments/battle/${battleId}`, data)
        return resp.data
    }
}

export const api = new ApiClient()
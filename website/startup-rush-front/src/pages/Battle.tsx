import { useEffect, useState } from "react"
import { api, BattleWithEvents } from "../services/ApiClient"
import { useParams } from "react-router-dom"
import { Header } from "../components/Header"
import { 
    Box, 
    Heading, 
    Grid, 
    Stack, 
    Text, 
    Button
} from "@chakra-ui/react"

export default function Battle() {
    const { id } = useParams<{ id: string }>()
    const [data, setData] = useState<BattleWithEvents | null>(null)
    const [selectionStartup1, setSelectionStartup1] = useState<number[]>([])
    const [selectionStartup2, setSelectionStartup2] = useState<number[]>([])
  
    useEffect(() => {
        api.getBattleById(id ?? "1")
        .then(response => {
            console.log('API Response:', response)
            setData(response.data)
        })
    }, [id])

    if (!data) {
        console.log('No data yet')
        return <div>Loading...</div>
    }

    const { battle, events } = data

    const toggleEvent = (eventId: number, startupNumber: 1 | 2) => {
        if (startupNumber === 1) {
            setSelectionStartup1(prev => 
                prev.includes(eventId) 
                    ? prev.filter(id => id !== eventId)
                    : [...prev, eventId]
            )
        } else {
            setSelectionStartup2(prev => 
                prev.includes(eventId) 
                    ? prev.filter(id => id !== eventId)
                    : [...prev, eventId]
            )
        }
    }

    const handleSubmitBattle = async () => {
        try {
            const battleData = {
                battle: [
                    { startupId: battle.startup1Id, eventIds: selectionStartup1 },
                    { startupId: battle.startup2Id, eventIds: selectionStartup2 }
                ]
            }
            
            await api.submitBattleEvents(id ?? "1", battleData)
        } catch (error) {
            console.error("Error submitting battle events:", error)
        }
    }

    if (battle.finished) {
        return (
            <div>
                <Header />
                <Box>
                    <Heading>Batalha #{battle.id} (Finalizada)</Heading>
                    <Text>
                        Startup {battle.startup1Id} vs Startup {battle.startup2Id}
                    </Text>
                    <Text>
                        Vencedor: Startup {battle.winnerId}
                    </Text>
                    <Text>
                        Placar: {battle.score1} x {battle.score2}
                    </Text>
                </Box>
            </div>
        )
    }

    return (
        <div>
            <Header />
            <Box>
                <Heading>Batalha #{battle.id}</Heading>
                
                <Grid templateColumns="1fr 1fr">
                    <Stack>
                        <Heading>Startup {battle.startup1Id}</Heading>
                        {events.map((event) => (
                            <Button
                                key={`s1-${event.id}`}
                                onClick={() => toggleEvent(event.id, 1)}
                                variant={selectionStartup1.includes(event.id) ? "solid" : "outline"}
                            >
                                {event.name} ({event.score > 0 ? '+' : ''}{event.score})
                            </Button>
                        ))}
                    </Stack>

                    <Stack>
                        <Heading>Startup {battle.startup2Id}</Heading>
                        {events.map((event) => (
                            <Button
                                key={`s2-${event.id}`}
                                onClick={() => toggleEvent(event.id, 2)}
                                variant={selectionStartup2.includes(event.id) ? "solid" : "outline"}
                            >
                                {event.name} ({event.score > 0 ? '+' : ''}{event.score})
                            </Button>
                        ))}
                    </Stack>
                </Grid>

                <Button
                    onClick={handleSubmitBattle}
                >
                    Finalizar Batalha
                </Button>
            </Box>
        </div>
    )
}
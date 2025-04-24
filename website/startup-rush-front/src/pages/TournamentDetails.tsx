import { useEffect, useState } from "react"
import { api, TournamentList } from "../services/ApiClient"
import { Header } from "../components/Header"
import { 
    Box,
    Heading, 
    SimpleGrid, 
    Text
} from "@chakra-ui/react"
import { useParams } from "react-router-dom"
import { BattleCard } from "../components/BattleCard"

export default function TournamentDetails() {
    const { id } = useParams<{ id: string }>()
    const [tournament, setTournament] = useState<TournamentList | null>(null)

    useEffect(() => {
        const loadTournamentData = async () => {
            try {
                const tournamentData = await api.getBattlesByTournament(id ?? '1')
                setTournament(tournamentData.data)
            } catch (err: any) {
                console.error('Error loading tournament data:', err)
            }
        }

        loadTournamentData()
    }, [id])

    return (
        <div>
            <Header />
            <Box p={4}>
                <Heading>
                    Torneio #{id} {tournament?.finished && <Text as="span" color="gray.500">(Finalizado)</Text>}
                </Heading>
                <SimpleGrid
                    columns={{ base: 1, md: 2, lg: 3 }}
                    gap={4}
                    justifyItems="center"
                >
                    {tournament && tournament.battles.length > 0 && tournament.battles.map((battle) => (
                        <BattleCard key={battle.id} battle={battle} />
                    ))}
                </SimpleGrid>
            </Box>
        </div>
    )
}
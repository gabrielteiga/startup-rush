import { useEffect, useState } from "react"
import { api, TournamentList } from "../services/ApiClient"
import { TournamentCard } from "../components/TournamentCard"
import { Header } from "../components/Header"
import { Flex, Heading, SimpleGrid, Text } from "@chakra-ui/react"

export default function HomePage() {
    const [tournaments, setTournaments] = useState<TournamentList[]>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        api.tournaments()
            .then((data) => setTournaments(data.data))
            .catch((err) => setError(err.message))
            .finally(() => setLoading(false))
        }, [])

        console.log(tournaments)

    return (
        <div>
            <Header />
            <div className="flex flex-col items-center min-h-screen bg-gray-100">
                <Heading>Meus Torneios</Heading>
                {loading && <Text >Carregando...</Text>}
                {error && <Text >{error}</Text>}
                <SimpleGrid
                  columns={{ base: 1, md: 2, lg: 3 }}
                  p={4}
                  gap={4}
                >
                    {tournaments.length > 0 && tournaments.map((tournament) => (
                        <TournamentCard key={tournament.id} tournament={tournament} />
                    ))}
                </SimpleGrid>
            </div>
        </div>
    )
}
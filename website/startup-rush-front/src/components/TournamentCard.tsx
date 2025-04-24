import React from 'react'
import { useNavigate } from 'react-router-dom'
import { TournamentList } from '../services/ApiClient'
import { Box, Heading, Text } from '@chakra-ui/react'

interface TournamentCardProps {
  tournament: TournamentList
}

export const TournamentCard: React.FC<TournamentCardProps> = ({ tournament }) => {
  const navigate = useNavigate()

  return (
    <Box
      onClick={() => navigate(`/torneios/${tournament.id}`)}
      cursor={'pointer'}
      border={'1px'}
      borderColor={'gray.200'}
      rounded={'md'}
      shadow={'md'}
      p={4}
      w={'300px'}
      _hover={{ shadow: 'lg',bgColor: 'gray.50' }}
    >
      <Heading>
        Torneio #{tournament.id}
      </Heading>

      <Text>
        Status:{' '}
        <span
          className={
            tournament.finished
              ? 'text-green-600 font-medium'
              : 'text-yellow-600 font-medium'
          }
        >
          {tournament.finished ? 'Finalizado' : 'Em andamento'}
        </span>
      </Text>

      <Text className="mb-1">
        Participantes:{' '}
        <span className="font-medium">{tournament.participants.length}</span>
      </Text>

      {tournament.battles && 
      <Text className="mb-4">
        Batalhas:{' '}
        <span className="font-medium">{tournament.battles.length}</span>
      </Text>}
    </Box>
  )
}
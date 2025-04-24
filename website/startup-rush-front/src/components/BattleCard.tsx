import React from 'react'
import { useNavigate } from 'react-router-dom'
import { Battle } from '../services/ApiClient'
import { Box, Heading, Text, Badge } from '@chakra-ui/react'

interface BattleCardProps {
  battle: Battle
}

const formatPhase = (phase: string) => {
  switch (phase) {
    case 'semi_final':
      return 'Semi Final'
    case 'final':
      return 'Final'
    default:
      return phase.charAt(0).toUpperCase() + phase.slice(1).replace('_', ' ')
  }
}

export const BattleCard: React.FC<BattleCardProps> = ({ battle }) => {
  const navigate = useNavigate()

  return (
    <Box
      onClick={() => navigate(`/torneios/battle/${battle.id}`)}
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
        Batalha #{battle.id}
      </Heading>

      <Badge mb={3} colorScheme={battle.phase === 'final' ? 'green' : 'blue'}>
        {formatPhase(battle.phase)}
      </Badge>

      <Text>
        Status:{' '}
        <span>
          {battle.finished ? 'Finalizado' : 'Em andamento'}
        </span>
      </Text>

      <Text>
        Participante 1 n°:{' '}
        <span className="font-medium">{battle.startup1Id}</span>
      </Text>
 
      <Text>
        Participante 2 n°:{' '}
        <span className="font-medium">{battle.startup2Id}</span>
      </Text>
    </Box>
  )
}
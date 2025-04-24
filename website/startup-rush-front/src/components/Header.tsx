import { Button, Flex, Heading, Link } from '@chakra-ui/react'
import React from 'react'

export const Header: React.FC = () => {
    return (
        <Flex
            as="header"
            justifyContent="space-between"
            alignItems="center"
            w="100%"
            p={10}
            marginBottom={4}
        >
            <Heading>StartupRush</Heading>
            <Flex gap={4}>
                <Link href="/">Início</Link>
                <Link href="/startups">Startups</Link>
                <Link href="/torneios">Torneios</Link>
            </Flex>
        </Flex>
    )
}
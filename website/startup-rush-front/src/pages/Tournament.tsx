import { useForm } from 'react-hook-form'
import { Header } from "../components/Header"
import { Button, Heading, Input, Stack, Text, CheckboxGroup, Checkbox } from "@chakra-ui/react"
import { useNavigate } from 'react-router-dom'
import { api, Startup } from '../services/ApiClient'
import { useEffect, useState } from 'react'

export default function Tournament() {
    const [startups, setStartups] = useState<Startup[]>([])
    const [selected, setSelected] = useState<number[]>([])
    const validCount = selected.length === 4 || selected.length === 8

    const onSubmit = async () => {
        const payload = {
            startupsIDs: selected
        }

        api.createTournament(payload).then((response) => {
            console.log(response)
            if (response === 201) {
                alert('Torneio cadastrado com sucesso')
            } else {
                alert('Erro ao cadastrar Torneio')
            }
        })
    }

    const handleToggle = (id: number, checked: boolean) => {
        setSelected(prev =>
          checked
            ? [...prev, id]
            : prev.filter(item => item !== id)
        )
      }
    
    useEffect(() => {
        api.listStartups()
          .then((resp) => setStartups(resp.data))
          .catch(err => alert(err.message))
      }, [])

    return (
        <div>
            <Header />
            <div className="flex flex-col items-center min-h-screen bg-gray-100 p-4">
                <Heading>Cadastro de Torneio</Heading>
                <div>
                    <Text>Selecione as Startups para o torneio (4 à 8)</Text>
                        <Stack>
                            {startups.map((startup) => (
                                <Checkbox.Root
                                    key={startup.id}
                                    value={startup.id.toString()}
                                    checked={selected.includes(startup.id)}
                                    onCheckedChange={(checked) => handleToggle(startup.id, checked.checked as boolean)}
                                >
                                    <Checkbox.HiddenInput/>
                                    <Checkbox.Control />
                                    <Checkbox.Label>
                                        {startup.name} - {startup.slogan}
                                    </Checkbox.Label>
                                </Checkbox.Root>
                            ))}
                        </Stack>

                    <Text>
                        É necessário selecionar 4 ou 8 startups para iniciar o torneio.
                    </Text>
                    <Button
                        type="submit"
                        onClick={onSubmit}
                        disabled={!validCount}
                        >Cadastrar</Button>
                </div>
            </div>
        </div>
    )
}
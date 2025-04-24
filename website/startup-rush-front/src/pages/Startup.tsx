import { useForm } from 'react-hook-form'
import { Header } from "../components/Header"
import { Button, Heading, Input, Stack, Text } from "@chakra-ui/react"
import { api } from '../services/ApiClient'

type FormRequest = {
    name: string
    slogan: string
    foundation: string
}

export default function Startup() {
    const { register, handleSubmit } = useForm<FormRequest>()
    const onSubmit = async (data: FormRequest) => {
        const payload = {
            name: data.name,
            slogan: data.slogan,
            foundation: new Date(data.foundation).toISOString(),
        }
        api.createStartup(payload).then((response) => {
            console.log(response)
            if (response === 201) {
                alert('Startup cadastrada com sucesso')
            } else {
                alert('Erro ao cadastrar startup')
            }
        })
    }
    
    return (
        <div>
            <Header />
            <div className="flex flex-col items-center min-h-screen bg-gray-100 p-4">
                <Heading>Cadastro de Startups</Heading>
                <form onSubmit={handleSubmit(onSubmit)}>
                    <Stack gap={4}>
                        <Text>Nome da Startup</Text>
                        <Input
                            typeof="text"
                            {...register("name", { required: true })}
                        />
                        <Text>Slogan</Text>
                        <Input
                            typeof="text"
                            {...register("slogan", { required: true })}
                        />
                        <Text>Ano de Fundação</Text>
                        <Input
                            type="date"
                            {...register("foundation", { required: true })}
                        />
                        <Button type="submit">Cadastrar</Button>
                    </Stack>
                </form>
            </div>
        </div>
    )
}
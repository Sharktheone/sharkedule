import {IconArrowBigLeft, IconError404, IconReload} from "@tabler/icons-react"
import {isRouteErrorResponse, Navigate, useNavigate, useRouteError} from "react-router-dom"
import {Button, Stack, useMantineTheme} from "@mantine/core"
import {Text} from "@/components/ui/ui"

import styles from "./styles.module.scss"
import {useEffect, useState} from "react"
import {Title} from "@/components/ui/title/title"

export default function BoardsError() {
    const error = useRouteError()
    const navigate = useNavigate()
    const theme = useMantineTheme()
    const [retry, setRetry] = useState(false)

    let intervalID: number

    function tryAgain() {
        if (document.visibilityState === "visible") {
            setRetry(true)
            setTimeout(
                () => {
                    navigate("")
                    setRetry(false)
                }, 1000)
        }
    }


    useEffect(() => {
        if (document.visibilityState === "visible") {
            //@ts-ignore
            intervalID = setInterval(() => {
                tryAgain()
            }, 10000) || intervalID
        } else {
            clearInterval(intervalID)
        }

    }, [document.visibilityState])

    useEffect(() => {
        return () => {
            clearInterval(intervalID)
        }
    }, [])


    function Navigation() {
        return (
            <Stack className={styles.errornav} align="center">
                <Button gradient={{from: 'teal', to: 'blue', deg: 60}} variant="gradient" onClick={() => navigate(-1)}>
                    <IconArrowBigLeft/>
                    <div>
                        Go Back
                    </div>
                </Button>
                <Button gradient={{from: "teal", to: "lime", deg: 105}} variant="gradient" onClick={tryAgain}>
                    <IconReload className={retry ? styles.rotate : undefined}/>
                    <div>
                        {retry ? "Retrying" : "Try Again"}
                    </div>
                </Button>
            </Stack>
        )
    }

    if (isRouteErrorResponse(error)) {
        switch (error.status) {
            case 401:
                return (
                    <Navigate to="/"/>
                )
            case 404:
                return (
                    <div>
                        <Title c="error">Board not found</Title>
                        <IconError404 size={"xl"}/>
                        <Navigation/>
                    </div>
                )
            case 500:
                return (
                    <div>
                        <Title c="error">Server Error</Title>
                        <Text s="medium" c="error">Please try again later</Text>
                        <Text s="medium" c="error"> {error.data.text} </Text>
                        <Navigation/>
                    </div>
                )
            case 503:
                return (
                    <div>
                        <Title>Server Error</Title>
                        <Text s="medium" c="error">
                            Looks like our API is down, please try again later
                        </Text>
                        <Text s="medium" c="error"> {error.data.text} </Text>
                        <Navigation/>
                    </div>
                )
        }

        return (
            <div>
                <Title s={1}>Error Loading Task Boards</Title>
                <Text s="medium" c="error">{error.data.text}</Text>
                <Navigation/>
            </div>
        )
    }
    return (
        <div>
            <Title c="error"> Error Loading Task Boards</Title>
            <Text c="error" s="medium">Unknown error</Text>
            <Navigation/>
        </div>
    )
}
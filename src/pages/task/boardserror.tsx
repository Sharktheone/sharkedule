import {
    IconArrowBigLeft,
    IconError404, IconReload
} from "@tabler/icons-react"
import {isRouteErrorResponse, Navigate, useNavigate, useRouteError} from "react-router-dom"
import {Button, Stack, Text, Title, useMantineTheme} from "@mantine/core"

import styles from "./styles.module.scss"

export default function BoardsError() {
    const error = useRouteError()
    const navigate = useNavigate()
    const theme = useMantineTheme()

    function Navigation() {
        return (
            <Stack className={styles.errornav} align="center">
                <Button gradient={{ from: 'teal', to: 'blue', deg: 60 }} variant="gradient" onClick={() => navigate(-1)}>
                    <IconArrowBigLeft/>
                    <div>
                        Go Back
                    </div>
                </Button>
                <Button gradient={{ from: "teal", to: "lime", deg: 105 }} variant="gradient" onClick={() => navigate("")}>
                    <IconReload/>
                    <div>
                        Try Again
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
                        <Title color={theme.colors.red[4]}>Board Not Found</Title>
                        <IconError404 size={"xl"}/>
                        <Navigation/>
                    </div>
                )
            case 500:
                return (
                    <div>
                        <Title color={theme.colors.red[4]}>Server Error</Title>
                        <Text color={theme.colors.red[4]} size={"lg"}>Please try again later</Text>
                        <Text color={theme.colors.red[4]}> {error.data.sorry} </Text>
                        <Navigation/>
                    </div>
                )
            case 503:
                return (
                    <div>
                        <Title color={theme.colors.red[4]}>Server Error</Title>
                        <Text color={theme.colors.red[4]} size={"lg"}>Looks like our API is down, please try again later</Text>
                        <Text color={theme.colors.red[4]}> {error.data.sorry} </Text>
                        <Navigation/>
                    </div>
                )
        }

        return (
            <div>
                <Title color={theme.colors.red[4]}>Error Loading Task Boards</Title>
                <Text color={theme.colors.red[4]} size={"lg"}>{error.data.sorry}</Text>
                <Navigation/>
            </div>
        )
    }
    return (
        <div>
            <Title color={theme.colors.red[4]}> Error Loading Task Boards</Title>
            <Text color={theme.colors.red[4]} pb="lg" size={"lg"}>Unknown error</Text>
            <Navigation/>
        </div>
    )
}
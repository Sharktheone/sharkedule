import {IconArrowBigLeft, IconError404, IconReload} from "@tabler/icons-react"
import {isRouteErrorResponse, Navigate, useNavigate, useRouteError} from "react-router-dom"
import {Button, Text, Title} from "@/components/ui"
import styles from "./styles.module.scss"
import {useEffect, useState} from "react"


export default function BoardsError() {
    const error = useRouteError()
    const navigate = useNavigate()
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
            <div className={styles.errornav}>
                <Button gradient variant="primary" onClick={() => navigate(-1)}>
                    <IconArrowBigLeft/>
                    <Text s="medium">
                        Go Back
                    </Text>
                </Button>
                <Button gradient variant="green" onClick={tryAgain}>
                    <IconReload className={retry ? styles.rotate : undefined}/>
                    <Text s="medium">
                        {retry ? "Retrying" : "Try Again"}
                    </Text>
                </Button>
            </div>
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
                    <div className={styles.error}>
                        <Title c="error">Board not found</Title>
                        <IconError404 size={"xl"}/>
                        <Navigation/>
                    </div>
                )
            case 500:
                return (
                    <div className={styles.error}>
                        <Title c="error">Server Error</Title>
                        <Text s="medium" c="error">Please try again later</Text>
                        <Text s="medium" c="error"> {error.data.text} </Text>
                        <Navigation/>
                    </div>
                )
            case 503:
                return (
                    <div className={styles.error}>
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
            <div className={styles.error}>
                <Title s={1}>Error Loading Task Boards</Title>
                <Text s="medium" c="error">{error.data.text}</Text>
                <Navigation/>
            </div>
        )
    }
    return (
        <div className={styles.error}>
            <Title c="error"> Error Loading Task Boards</Title>
            <Text c="error" s="medium">Unknown error</Text>
            <Navigation/>
        </div>
    )
}
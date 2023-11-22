import {MenuView} from "@/components/menu/types"
import {useMemo, useState} from "react"
import styles from "./styles.module.scss"
import RenderView from "@/components/menu/RenderView"


type Props = {
    views: MenuView[]
    defaultView: string
}


export default function Menu_old({views, defaultView}: Props) {
    const [currentView, setCurrentView] = useState<string>(defaultView)
    const [secondsToReturn, setSecondsToReturn] = useState<number>(5)

    function returnToDefault() {
        setCurrentView(defaultView)
    }

    function secondsToDefault() {
        setTimeout(() => {
            if (secondsToReturn === 0) {
                returnToDefault()
                return
            }
            setSecondsToReturn(secondsToReturn - 1)
        }, 1000)
        return null
    }

    function getCurrentView() {
        const view = useMemo(() => views?.find((v) => v?.id === currentView), [currentView, views])
        if (!view) {
            return {
                id: "error",
                name: "Error, view not found",
                items: (
                    <div>
                        Error, view not found
                        <span> Returning to default in {secondsToReturn}s {secondsToDefault()} </span>
                    </div>
                )
            } as MenuView
        }
        return view
    }

    return (
        <div className={`${styles.menu} ${cx(classes.menu)}`}>
            <RenderView view={getCurrentView()}/>
        </div>
    )
}
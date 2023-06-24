import styles from "./styles.module.scss"
import {CSSProperties, useEffect, useRef, useState} from "react"
import {useColors} from "./styles"
import {Entries} from "@kanban/contextmenu/entries"

type Props = {
    x: number
    y: number
    open: boolean
    close: () => void
    type: types
}

export enum types {
    TASK,
    COLUMN
}

export default function ContextMenu({x, y, open, close}: Props) {
    const {classes, cx} = useColors()
    const [opened, setOpened] = useState(false)
    const [freshlyClosed, setFreshlyClosed] = useState(false)
    const [hidden, setHidden] = useState(true)
    const contextMenuRef = useRef<HTMLDivElement>(null)


    function position(): CSSProperties {
        return {
            top: `${y}px`,
            left: `${x}px`,
        }
    }

    useEffect(() => {
        if (open) {
            setHidden(false)
            setTimeout(() => {
                setOpened(true)
            }, 500)
            console.log("freshly opened")
            contextMenuRef?.current?.focus()
        } else {
            setFreshlyClosed(true)

            setOpened(false)
            setTimeout(() => {
                setFreshlyClosed(false)
                setHidden(true)
            }, 500)
        }
        console.log("open", open, "hidden", hidden, "freshlyOpened", opened, "freshlyClosed", freshlyClosed)
    }, [open])


    if (hidden) return null


    return (
        <div ref={contextMenuRef}
             className={`${styles.contextmenu} ${cx(classes.contextMenu)} ${opened ? styles.opened : ""} ${freshlyClosed ? styles.closed : ""}`}
             style={position()}>
            <div>
            <button onClick={close}/>
            {
                Entries.map((entry, i) => (
                    <div className={styles.entry} key={i}>
                        <div className={styles.icon}>
                            {entry.icon}
                        </div>
                        <div className={styles.name}>
                            {entry.name}
                        </div>
                    </div>

                ))
            }
            </div>
        </div>
    )
}
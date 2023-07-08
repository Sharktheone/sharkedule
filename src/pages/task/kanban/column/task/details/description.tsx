import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {useColors} from "./colors"
import styles from "./styles.module.scss"

type Props = {
    uuid: string
}

export  default function Description({uuid}: Props) {
    const {environment} = useContext(EnvironmentContext)

    const task = environment?.tasks?.find((task) => task.uuid === uuid)

    const {cx, classes} = useColors()

    if (!task) return null


    return (
        <textarea className={`${cx(classes.description)} ${styles.description}`}>
            {task.description}
        </textarea>
    )

}
import {ReactNode} from "react"
import {IconPlus, IconTrash} from "@tabler/icons-react"


export type ContextMenuEntry = {
    name: string
    color?: string
    handler?: (uuid: string) => void
    icon: ReactNode
    type: types
}

export enum types {
    TASK,
    COLUMN,
    BOTH
}


export const Entries: ContextMenuEntry[] = [
    {
        name: "Add Task",
        icon: <IconPlus/>,
        type: types.COLUMN,
    },
    {
        name: "Delete",
        color: "red",
        icon: <IconTrash/>,
        type: types.BOTH,
    }
]
import {types} from "@kanban/contextmenu/contextmenu"
import {ReactNode} from "react"
import {IconPlus, IconTrash} from "@tabler/icons-react"


export type ContextMenuEntry = {
    name: string
    color?: string
    // handler?: (uuid: string) => void
    icon: ReactNode
    // type: types
}

export const Entries: ContextMenuEntry[] = [
    {
        name: "Add Task",
        icon: <IconPlus/>,
    },
    {
        name: "Delete",
        color: "red",
        icon: <IconTrash/>,
    }
]
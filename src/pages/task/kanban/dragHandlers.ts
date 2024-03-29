import {ghostType, useGhost} from "@/pages/task/kanban/ghost"
import {DragStart, DragUpdate, DropResult} from "@hello-pangea/dnd"
import {Dispatch, SetStateAction} from "react"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"
import {Board, environment} from "@kanban/types"
import {toast} from "react-toastify"


export class dragHandlers {
    public readonly ghost: ghostType | undefined
    private readonly addGhost: (event: DragStart) => void
    private readonly removeGhost: () => void
    private readonly updateGhost: (event: DragUpdate) => void
    private readonly environment: environment
    private readonly setEnvironment: Dispatch<SetStateAction<environment>>
    private readonly navigate: ReturnType<typeof useNavigate>
    private readonly uuid: string

    constructor(environment: environment, setEnvironment: Dispatch<SetStateAction<environment>>, uuid: string) {
        const {ghost, addGhost, removeGhost, updateGhost} = useGhost()
        this.navigate = useNavigate()

        this.ghost = ghost
        this.addGhost = addGhost.bind(this)
        this.removeGhost = removeGhost.bind(this)
        this.updateGhost = updateGhost.bind(this)
        this.environment = environment
        this.setEnvironment = setEnvironment
        this.uuid = uuid
    }

    start(event: DragStart) {
        this.addGhost(event)
    }

    update(event: DragUpdate) {
        this.updateGhost(event)
    }

    end(event: DropResult) {
        if (event.type === "task") {
            this.removeGhost()

            let {destination, source, draggableId} = event
            if (!destination) return
            if (destination.droppableId === source.droppableId && destination.index === source.index) return
            this.reorderTask(source.droppableId, draggableId, destination.index, destination.droppableId)
        } else if (event.type === "column") {
            let {destination, source, draggableId} = event
            if (!destination) return
            if (destination.index === source.index) return
            this.reorderColumn(draggableId, destination.index)

        }
    }

    private reorderColumn(uuid: string, to: number) {
        let newBoard = {...this?.environment?.boards?.find((board) => board.uuid === this?.uuid)}

        let columnIndex = newBoard?.columns?.findIndex((column) => column === uuid)
        if (columnIndex === undefined) return
        let [column] = newBoard?.columns?.splice(columnIndex, 1) ?? []
        newBoard?.columns?.splice(to, 0, column)
        this?.setEnvironment({...this?.environment, boards: [newBoard] as Board[]})
        api.patch(`/${this.environment.workspace}/kanban/board/${this?.uuid}/column/${uuid}/move`, {
            index: to
        }).then((res) => {
            if (res.status > 300) {
                toast(`Error: ${res.data}`, {type: "error"})
                console.log(res)
            }
            this?.refresh()

        }).catch((err) => {
            toast(`Error: ${err.message}`, {type: "error"})
            console.log(err)
            this?.refresh()
        })
    }

    private reorderTask(fromColumn: string, uuid: string, to: number, toColumn: string,) {
        let columns = [...this.environment?.columns]
        if (!columns) return

        let fromTasks = columns?.find((column) => column.uuid === fromColumn)?.tasks
        if (!fromTasks) return

        let taskIndex = fromTasks?.findIndex((task) => task === uuid)
        let [task] = fromTasks?.splice(taskIndex, 1)


        let toTasks = columns?.find((column) => column.uuid === toColumn)?.tasks

        toTasks?.splice(to, 0, task)
        this.setEnvironment(
            {
                ...this.environment,
                columns: columns
            }
        )

        api.patch(`/${this.environment.workspace}/kanban/board/${this?.uuid}/column/${fromColumn}/task/${uuid}/move`, {
            column: toColumn,
            index: to
        }).then((res) => {
            if (res.status > 300) {
                toast(`Error: ${res.data}`, {type: "error"})
                console.log(res)
            }
            this.refresh()

        }).catch((err) => {
            toast(`Error: ${err.message}`, {type: "error"})
            console.log(err)
            this.refresh()
        })
    }

    private refresh() {
        this.navigate("")
    }
}
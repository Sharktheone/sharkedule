import {ghostType, useGhost} from "@/pages/task/kanban/ghost"
import {DragStart, DragUpdate, DropResult} from "react-beautiful-dnd"
import {Dispatch, SetStateAction} from "react"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {useNavigate} from "react-router-dom"
import {Board, environment} from "@kanban/types2"


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

    Start(event: DragStart) {
        this.addGhost(event)
    }

    Update(event: DragUpdate) {
        this.updateGhost(event)
    }

    End(event: DropResult) {
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
        let newBoard = {...this?.environment.boards?.find((board) => board.uuid === this?.uuid)}

        let columnIndex = newBoard?.columns?.findIndex((column) => column === uuid)
        if (columnIndex === undefined) return
        let [column] = newBoard?.columns?.splice(columnIndex, 1) ?? []
        newBoard?.columns?.splice(to, 0, column)
        this?.setEnvironment({...this?.environment, boards: [newBoard] as Board[]})
        api.patch(`/kanban/board/${this?.uuid}/column/${uuid}/move`, {
            index: to
        }).then((res) => {
            if (res.status > 300) {
                notifications.show({title: "Error", message: res.data, color: "red"})
                console.log(res)
            }
            this?.refresh()

        }).catch((err) => {
            notifications.show({title: "Error", message: err.message, color: "red"})
            console.log(err)
            this?.refresh()
        })
    }

    private reorderTask(fromColumn: string, uuid: string, to: number, toColumn: string,) {
        let newBoard = {...this?.environment.boards?.find((board) => board.uuid === this?.uuid)}
        if (!newBoard) return

        let fromTasks = this.environment.columns?.find((column) => column.uuid === fromColumn)?.tasks
        if (!fromTasks) return

        let taskIndex = fromTasks?.findIndex((task) => task === uuid)
        let [task] = fromTasks?.splice(taskIndex, 1)


        let toTasks = this.environment.columns?.find((column) => column.uuid === toColumn)?.tasks

        toTasks?.splice(to, 0, task)
        this.setEnvironment(
            {
                ...this.environment,
                columns: {
                    ...this.environment.columns,
                }
            }
        )

        console.log(this.environment)

        api.patch(`/kanban/board/${this?.uuid}/column/${fromColumn}/task/${uuid}/move`, {
            column: toColumn,
            index: to
        }).then((res) => {
            if (res.status > 300) {
                notifications.show({title: "Error", message: res.data, color: "red"})
                console.log(res)
            }
            this.refresh()

        }).catch((err) => {
            notifications.show({title: "Error", message: err.message, color: "red"})
            console.log(err)
            this.refresh()
        })
    }

    private refresh() {
        this.navigate("")
    }
}
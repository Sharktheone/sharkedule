import {Group} from '@mantine/core'
import {DragDropContext, DropResult} from "react-beautiful-dnd"
import Column from "./column/column"
import {useState} from "react"
import {kanbanBoardType} from "./types"
import {useLoaderData} from "react-router-dom"

export default function Kanban() {
    const loaderData = useLoaderData()
    const [board, setBoard] = useState<kanbanBoardType>(loaderData as kanbanBoardType)

    function dragEndHandler(result: DropResult) {
        let {destination, source, draggableId} = result
        console.log(result)
        if (!destination) return
        if (destination.droppableId === source.droppableId && destination.index === source.index) return
        reorderTask(source.droppableId, draggableId, destination.index, destination.droppableId)
    }

    function reorderTask(fromColumn: string, uuid: string, to: number, toColumn: string,) {
        let newBoard = {...board}

        let fromColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === fromColumn)
        let toColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === toColumn)
        let taskIndex = newBoard?.columns[fromColumnIndex]?.tasks?.findIndex((task) => task.uuid === uuid)
        let [task] = newBoard?.columns[fromColumnIndex]?.tasks?.splice(taskIndex, 1)


        newBoard?.columns[toColumnIndex]?.tasks?.splice(to, 0, task)
        setBoard(newBoard)
    }

    function renameTask(uuid: string, name: string) {
        let newBoard = {...board}
        newBoard.columns?.forEach((column) => {
            column.tasks?.forEach((task) => {
                if (task.uuid === uuid) {
                    task.name = name
                    return
                }
            })
        })
        setBoard(newBoard)
    }

    function renameColumn(uuid: string, name: string) {
        let newBoard = {...board}
        newBoard.columns?.forEach((column) => {
            if (column.uuid === uuid) {
                column.name = name
                return
            }
        })
        setBoard(newBoard)
    }


    return (
        <DragDropContext onDragEnd={dragEndHandler}>
            <Group position="center" align="start" noWrap={true}>
                {board.columns?.map((column) => (
                    <Column key={column.uuid} column={column} renameColumn={renameColumn} renameTask={renameTask}/>
                ))}
            </Group>
        </DragDropContext>
    )
}


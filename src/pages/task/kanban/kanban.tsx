import {Group} from '@mantine/core';
import test_data from "./test_data.json"
import {DragDropContext, DropResult, ResponderProvided} from "react-beautiful-dnd"
import Column from "./column/column";
import {useState} from "react"
import {kanbanBoardType} from "./types"
export default function Kanban() {
    const [board, setBoard] = useState<kanbanBoardType>(test_data as unknown as kanbanBoardType)

    function dragEndHandler(result: DropResult, provided: ResponderProvided) {
        let {destination, source, draggableId} = result
        console.log(result)
        if (!destination) return
        if (destination.droppableId === source.droppableId && destination.index === source.index) return
        reorderTask(source.droppableId, draggableId, destination.index, destination.droppableId)
    }

    function reorderTask(fromColumn: string, uuid: string, to : number, toColumn: string, ) {
        let newBoard = {...board}

        let fromColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === fromColumn)
        let toColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === toColumn)
        let taskIndex = newBoard?.columns[fromColumnIndex]?.tasks?.findIndex((task) => task.uuid === uuid)
        let [task] = newBoard?.columns[fromColumnIndex]?.tasks?.splice(taskIndex, 1)


        newBoard?.columns[toColumnIndex]?.tasks?.splice(to, 0, task)
        setBoard(newBoard)
    }


    return (
        <DragDropContext onDragEnd={dragEndHandler}>
            <Group position="center" align="start" noWrap={true}>
                {board.columns?.map((column) => (
                    <Column key={column.uuid} column={column}/>
                ))}
            </Group>
        </DragDropContext>
)
}


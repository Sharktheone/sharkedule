import {DragStart, DragUpdate} from "react-beautiful-dnd"
import {useState} from "react"

export type ghostType = {
    index: number
    height: number
    hoveredColumnID: string

}

export type ghostElementType = {
    offsetTop: string
    height: string
}

export function getDraggedElement(draggableId: string) {
    return document.querySelector(`[data-rbd-drag-handle-draggable-id='${draggableId}'] > div > div`)
}


export function useGhost() {
    const [ghost, setGhost] = useState<ghostType | undefined>(undefined)

    function addGhost(event: DragStart) {
        if (event.type === "task") {
            let draggedElement = getDraggedElement(event.draggableId)
            if (!draggedElement) return
            console.log(draggedElement)

            let rect = draggedElement.getBoundingClientRect()

            setGhost({
                height: rect.height,
                index: event.source.index,
                hoveredColumnID: event.source.droppableId,
            })
        } else if (event.type === "column") {

        }

        console.log("drag start")
    }

    function updateGhost(event: DragUpdate) {
        if (event.type === "task") {
            if (event.destination == null) {
                setGhost(undefined)
                return
            }

            let draggedElement = getDraggedElement(event.draggableId)
            if (!draggedElement) return

            let rect = draggedElement.getBoundingClientRect()

            setGhost({
                height: rect.height,
                index: event.destination?.index ?? ghost?.index ?? 0,
                hoveredColumnID: event.destination?.droppableId ?? ghost?.hoveredColumnID ?? "",
            })
        } else if (event.type === "column") {

        }
    }

    function removeGhost() {
        setGhost(undefined)
    }

    return {
        ghost,
        addGhost,
        removeGhost,
        updateGhost,
    }
}
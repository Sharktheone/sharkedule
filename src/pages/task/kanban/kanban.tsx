import {Button, CloseButton, Input, Stack, Text, Title} from '@mantine/core'
import {DragDropContext, Droppable} from "react-beautiful-dnd"
import Column from "./column/column"
import {useEffect, useRef, useState} from "react"
import {kanbanBoardType} from "./types"
import {useLoaderData} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {useStyles} from "./styles"
import {dragHandlers} from "./dragHandlers"
import {handlers} from './handlers'

export default function Kanban() {
    const loaderData = useLoaderData()
    const [board, setBoard] = useState<kanbanBoardType>(loaderData as kanbanBoardType)
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const {classes, cx} = useStyles()

    const drag = new dragHandlers(board, setBoard)
    const h = new handlers(board, setBoard, setIsAdding, newColRef)

    useEffect(() => {
        setBoard(loaderData as kanbanBoardType)
    }, [loaderData])

    useEffect(() => {
        newColRef?.current?.focus()
    }, [isAdding])


    return (
        <div className={styles.board}>
            <Title order={1} align="center">{board.name}</Title>
            <Text mb="sm" align="center" color="dimmed">Drag and drop tasks to reorder them</Text>
            <DragDropContext onDragStart={event => drag.Start(event)} onDragEnd={event => drag.End(event)}
                             onDragUpdate={event => drag.Update(event)}>
                <Droppable droppableId={board.uuid} type="column" direction="horizontal">
                    {(provided) => (
                        <div
                            ref={provided.innerRef}
                            {...provided.droppableProps}>
                            <div className={styles.cols}>
                                {board.columns?.map((column) => (
                                    <div key={column.uuid}>
                                        <Column column={column}
                                                setBoard={setBoard}
                                                board={board}
                                                index={board.columns?.indexOf(column) ?? 0}
                                                boardUUID={board.uuid} ghost={drag.ghost}/>
                                    </div>
                                ))}


                                {provided.placeholder}

                                {!isAdding ?
                                    <>
                                        <button onClick={() => h.handleNewColumn()}
                                                className={`${cx(classes.addColumn)} ${styles.footer}`}>
                                            <IconPlus size={24}/>
                                            <Text align="center">Add a Column</Text>
                                        </button>
                                    </> :
                                    <div>
                                        <Stack className={styles.add}>
                                            <Input ref={newColRef} onBlur={() => h.cancelAddColumn()}
                                                   placeholder="Column name"/>
                                            <div className={styles.menu}>
                                                <Button onClick={() => h.addColumn()}
                                                        gradient={{from: "#6dd6ed", to: "#586bed"}}
                                                        variant="gradient">Create
                                                </Button>
                                                <CloseButton onClick={() => setIsAdding(false)}/>
                                            </div>
                                        </Stack>
                                    </div>

                                }

                            </div>
                        </div>
                    )

                    }

                </Droppable>

            </DragDropContext>
        </div>
    )
}


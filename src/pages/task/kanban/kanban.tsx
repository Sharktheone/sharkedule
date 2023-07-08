import {Button, CloseButton, Input, Stack, Text, Title} from '@mantine/core'
import {DragDropContext, Droppable} from "react-beautiful-dnd"
import Column from "@/pages/task/kanban/column/column"
import {useEffect, useMemo, useRef, useState} from "react"
import {useLoaderData, useNavigate, useParams} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {useStyles} from "./styles"
import {dragHandlers} from "@/pages/task/kanban/dragHandlers"
import {handlers} from '@/pages/task/kanban/handlers'
import {environment} from "@kanban/types"
import {EnvironmentProvider} from "@kanban/environment"
import Tstyles from "@kanban/column/task/styles.module.scss"
import ContextMenu from "@kanban/contextmenu/contextmenu"

export default function Kanban() {
    const loaderData = useLoaderData()
    const uuid = useParams().uuid
    const [environment, setEnvironment] = useState<environment>(loaderData as environment)
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const boardRef = useRef<HTMLDivElement>(null)
    const [contextMenu, setContextMenu] = useState({open: false, x: 0, y: 0})
    const {classes, cx} = useStyles()

    const navigate = useNavigate()

    if (uuid === undefined) {
        navigate("../")
        return null
    }

    useEffect(() => {
        boardRef?.current?.addEventListener("contextmenu", (e) => {
            // if (contextMenu.open) return
            e.preventDefault()
            setContextMenu({
                open: true,
                x: e.pageX,
                y: e.pageY
            })

            if (e.target == null) return
            if (!("parentElement" in e.target)) return
            let element: HTMLElement = e.target.parentElement as HTMLElement

            for (let i = 0; i < 99; i++) {
                if (element.classList.contains(Tstyles.task)) {
                    console.warn("found task!!!")
                    break
                }
                const e = element.parentElement
                if (e == null) break
                element = e
            }
        })
    }, [])


    const drag = new dragHandlers(environment, setEnvironment, uuid)
    const h = new handlers(setIsAdding, newColRef, uuid)

    useEffect(() => {
        setEnvironment(loaderData as environment)
    }, [loaderData])

    useEffect(() => {
        newColRef?.current?.focus()
    }, [isAdding])

    function getBoard() {
        return useMemo(() => {
            return environment?.boards?.find(b => b.uuid === uuid)
        }, [environment])
    }

    function closeContextMenu() {
        setContextMenu({...contextMenu, open: false})
    }


    return (
        <EnvironmentProvider environment={environment} setEnvironment={setEnvironment}>
            <div className={styles.board} ref={boardRef}>
                <Title order={1} align="center">{getBoard()?.name}</Title>
                <Text mb="sm" align="center" color="dimmed">Drag and drop tasks to reorder them</Text>
                <DragDropContext onDragStart={event => drag.Start(event)} onDragEnd={event => drag.End(event)}
                                 onDragUpdate={event => drag.Update(event)}>
                    <Droppable droppableId={uuid} type="column" direction="horizontal">
                        {(provided) => (
                            <div
                                ref={provided.innerRef}
                                {...provided.droppableProps}>
                                <div className={styles.cols}>
                                    {getBoard()?.columns?.map((column) => (
                                        <div key={column}>
                                            <Column column={column}
                                                    boardUUID={uuid} ghost={drag.ghost}/>
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
                                                       onKeyDown={(e) => {
                                                           if (e.key === "Enter") {
                                                               h.addColumn()
                                                           }
                                                       }}
                                                       onSubmit={() => h.addColumn()}
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
            <ContextMenu x={contextMenu.x} y={contextMenu.y} open={contextMenu.open}
                         close={closeContextMenu}/>
        </EnvironmentProvider>
    )
}


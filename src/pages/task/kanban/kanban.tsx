import {Button, CloseButton, Input, Text, Title} from "@/components/ui"
import {DragDropContext, Droppable} from "@hello-pangea/dnd"
import Column from "@/pages/task/kanban/column/column"
import {useEffect, useMemo, useRef, useState} from "react"
import {useLoaderData, useNavigate, useParams} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {dragHandlers} from "@/pages/task/kanban/dragHandlers"
import {handlers} from '@kanban/handlers'
import {environment} from "@kanban/types"
import {EnvironmentProvider} from "@kanban/environment"
import Tstyles from "@kanban/column/task/styles.module.scss"
import ContextMenu from "@kanban/contextmenu/contextmenu"
import {ColorSelector} from "@/components/colorselector/colorselector"


export default function Kanban() {
    const loaderData = useLoaderData()
    const board = useParams().board
    const workspace = useParams().workspace
    const [environment, setEnvironment] = useState<environment>(loaderData as environment)
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const boardRef = useRef<HTMLDivElement>(null)
    const [contextMenu, setContextMenu] = useState({open: false, x: 0, y: 0})

    const navigate = useNavigate()

    if (!workspace) {
        navigate("../")
        return null
    }

    if (!board) {
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

    useEffect(() => {
        console.log(environment)
    }, [environment])


    const drag = new dragHandlers(environment, setEnvironment, board)
    const h = new handlers(setIsAdding, newColRef, board, environment.workspace)

    useEffect(() => {
        setEnvironment(loaderData as environment)
    }, [loaderData])

    useEffect(() => {
        newColRef?.current?.focus()
    }, [isAdding])

    function getBoard() {
        return useMemo(() => {
            return environment?.boards?.find(b => b.uuid === board)
        }, [environment])
    }

    function closeContextMenu() {
        setContextMenu({...contextMenu, open: false})
    }


    return (
        <EnvironmentProvider environment={environment} setEnvironment={setEnvironment}>
            <ColorSelector/>
            <div className={styles.board} ref={boardRef}>
                <Title className={styles.title} s={1} a="left">{getBoard()?.name}</Title>
                <Text a="left" color="dimmed">Drag and drop tasks to reorder them</Text>
                <DragDropContext onDragStart={event => drag.Start(event)} onDragEnd={event => drag.End(event)}
                                 onDragUpdate={event => drag.Update(event)}>
                    <Droppable droppableId={board} type="column" direction="horizontal">
                        {(provided) => (
                            <div
                                ref={provided.innerRef}
                                {...provided.droppableProps}>
                                <div className={styles.cols}>
                                    {getBoard()?.columns?.map((column) => (
                                        <div key={column}>
                                            <Column column={column}
                                                    boardUUID={board} ghost={drag.ghost}/>
                                        </div>
                                    ))}


                                    {provided.placeholder}

                                    {!isAdding ?
                                        <>
                                            <button onClick={() => h.handleNewColumn()}
                                                    className={styles.footer}>
                                                <IconPlus size={24}/>
                                                <Text a="center">Add a Column</Text>
                                            </button>
                                        </> :
                                        <div>
                                            <div className={styles.add}>
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
                                                            gradient
                                                            variant="primary">Create
                                                    </Button>
                                                    <CloseButton onClick={() => setIsAdding(false)}/>
                                                </div>
                                            </div>
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


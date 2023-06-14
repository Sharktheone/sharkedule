import {Button, CloseButton, Input, Stack, Text, Title} from '@mantine/core'
import {DragDropContext, Droppable} from "react-beautiful-dnd"
import Column from "./column/column"
import {useEffect, useMemo, useRef, useState} from "react"
import {useLoaderData, useNavigate, useParams} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {useStyles} from "./styles"
import {dragHandlers} from "./dragHandlers"
import {handlers} from './handlers'
import {environment} from "@kanban/types2"
import {EnvironmentProvider} from "@kanban/environment"

export default function Kanban() {
    const loaderData = useLoaderData()
    const uuid = useParams().uuid
    const [environment, setEnvironment] = useState<environment>(loaderData as environment)
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const {classes, cx} = useStyles()

    const navigate = useNavigate()

    if (uuid === undefined) {
        navigate("../")
        return
    }

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
            return environment.boards?.find(b => b.uuid === uuid)

        }, [environment])
    }

    return (
        <EnvironmentProvider environment={environment} setEnvironment={setEnvironment}>
            <div className={styles.board}>
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
                                                    setEnvironment={setEnvironment}
                                                    environment={environment}
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
        </EnvironmentProvider>
    )
}


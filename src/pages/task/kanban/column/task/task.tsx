import {Text} from "@mantine/core"
import {useStyles} from "./styles"
import React, {RefObject, useContext, useRef, useState} from "react"
import styles from "./styles.module.scss"
import {IconCircleCheck, IconTrash, IconX} from "@tabler/icons-react"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {useNavigate} from "react-router-dom"
import {SlotProvider} from "@kanban/column/task/slots/slotProvider"
import UpperSlot from "@kanban/column/task/slots/upper/upperSlot"
import LowerSlot from "@kanban/column/task/slots/lower/lowerSlot"
import {EnvironmentContext} from "@kanban/environment"
import TaskDetails from "@kanban/column/task/details/details"
import useDoubleClick from "@/hooks/useDoubleClick/useDoubleClick"

type TaskProps = {
    board: string
    column: string
    task: string
    renameTask: (uuid: string, name: string) => void
}


export default function Task({task, renameTask, board, column}: TaskProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)
    const navigate = useNavigate()
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const [taskDetails, setTaskDetails] = useState(false)

    const clickHandler = useDoubleClick(openDetails, editText)

    const nameRef = useRef<HTMLDivElement>()


    const [t, setT] = useState(getTask(task))


    function getTask(uuid: string) {
        return environment?.tasks?.find((task) => task.uuid === uuid)
    }

    function editText() {
        setEditable(true)
        nameRef?.current?.focus()

    }

    function handleBlur(e: any) {
        setEditable(false)
        renameTask(task, e.target.innerText)
    }

    function handleDelete() {
        api.delete(`/kanban/board/${board}/column/${column}/task/${task}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    notifications.show({title: "Success", message: "Deleted Task", color: "green"})
                    navigate("#")
                }
            }
        ).catch(
            (err) => {
                notifications.show({title: "Error", message: err.message, color: "red", icon: <IconX/>})
                console.log(err)
            }
        )
    }


    function openDetails() {
        if (editable) return
        setTaskDetails(true)
    }


    return (
        <SlotProvider task={task}>
            <div className={`${cx(classes.task)} ${styles.task}`} onClick={clickHandler.onClick}
                 onDoubleClick={clickHandler.onDoubleClick}>
                <UpperSlot/>
                <div className={styles.taskname}>
                    <div className={styles.name}>
                        <IconCircleCheck/>
                        <Text align="start" onBlur={handleBlur} contentEditable={editable}
                              ref={nameRef as RefObject<HTMLDivElement>}> {/*TODO: dont use contentEditable*/}
                            {t?.name as string}
                        </Text>
                    </div>
                    <div className={styles.hover}>
                        <div>
                            <button onClick={handleDelete}>
                                <IconTrash/>
                            </button>
                        </div>
                    </div>
                </div>
                <LowerSlot/>
            </div>
            <TaskDetails open={taskDetails} setOpen={setTaskDetails} uuid={task}/>
        </SlotProvider>
    )
}
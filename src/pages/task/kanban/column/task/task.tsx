import {Text} from "@mantine/core"
import {useStyles} from "./styles"
import {Dispatch, SetStateAction, useState} from "react"
import styles from "./styles.module.scss"
import {IconCircleCheck, IconTrash, IconX} from "@tabler/icons-react"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {useNavigate} from "react-router-dom"
import {SlotProvider} from "@kanban/column/task/slots/slotProvider"
import UpperSlot from "@kanban/column/task/slots/upper/upperSlot"
import LowerSlot from "@kanban/column/task/slots/lower/lowerSlot"
import {environment} from "@kanban/types2"

type TaskProps = {
    board: string
    column: string
    task: string
    environment: environment
    setEnvironment: Dispatch<SetStateAction<environment>>
    renameTask: (uuid: string, name: string) => void
}


export default function Task({task, renameTask, board, column, environment, setEnvironment}: TaskProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)
    const navigate = useNavigate()

    function editText() {
        setEditable(true)
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

    return (
        <SlotProvider task={task} environment={environment} setEnvironment{setEnvironment}>
            <div className={`${cx(classes.task)} ${styles.task}`}>
                <UpperSlot/>
                <div className={styles.taskname}>
                    <div className={styles.name}>
                        <IconCircleCheck/>
                        <Text align="start" onClick={editText} onBlur={handleBlur} contentEditable={editable}>
                            {task.name}
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
        </SlotProvider>
    )
}
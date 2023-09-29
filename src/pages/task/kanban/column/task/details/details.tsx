import Details from "@/components/details/details"
import {Dispatch, SetStateAction, useState} from "react"
import {Task} from "@kanban/types"
import {getTask} from "@/pages/task/utils/task"
import TagSelector from "@/components/kanban/tags/selector/selector"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"
import {Menu} from "@/components/menu/menu"
import Description from "@kanban/column/task/details/description"
import styles from "./styles.module.scss"

type Props = {
    open: boolean
    setOpen: Dispatch<SetStateAction<boolean>>
    uuid: string
}

export default function TaskDetails({open, setOpen, uuid}: Props) {
    const [task, setTask] = useState<Task>(getTask(uuid) ?? {} as Task)
    const navigate = useNavigate()

    function onClose() {
        setOpen(false)
    }

    function setTags(tags: string[]) {
        api.patch(`/kanban/task/${uuid}/tags`, {tags: tags}).then(
            (res) => {
                if (res.status >= 300) {
                    console.log(res.data)
                } else {
                    navigate("")
                }
            }).catch(
            (err) => {
                console.log(err)
            })
    }

    return (
        <Details open={open} onClose={onClose} title={task.name}>
            <div className={styles.details}>
                <Description uuid={uuid}/>
                <Menu defaultView="anotherView">
                    <Menu.View id="default" name="Edit">
                        <div>
                            HELLO YOU LITTLE THING
                            LKSJFLGKJSFLKGJSFLKGJSLFKGJSLFKGJLSKFJGLSKFJGLSKFJGLÖSKJGÖLSKGJ
                        </div>
                    </Menu.View>
                    <Menu.View id="anotherView" name="Edit">
                        <div>
                            ABCDEFGHIJKLMNOPQRSTUVWXYZ
                            LKSJLGKJFSLKGJLSKFJGSLFKJGLSKFJGLKJGLKDSFJGLKDJF
                        </div>
                    </Menu.View>
                </Menu>
                <TagSelector selected={task.tags} onChange={(tags) => setTags(tags)}/>
            </div>
        </Details>
    )
}
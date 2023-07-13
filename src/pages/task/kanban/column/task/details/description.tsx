import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {useColors} from "./colors"
import styles from "./styles.module.scss"
import LoaderOverlay from "@/components/loaderOverlay/loaderOverlay"
import {useDebouncedState} from "@mantine/hooks"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"
import {notifications} from "@mantine/notifications"
import {useEffect} from "react"

type Props = {
    uuid: string
}

export default function Description({uuid}: Props) {
    const {environment} = useContext(EnvironmentContext)

    const task = environment?.tasks?.find((task) => task.uuid === uuid)

    const [description, setDescription] = useDebouncedState<string>(task?.description ?? "", 1000)

    const {cx, classes} = useColors()

    const navigate = useNavigate()

    if (!task) return null

    useEffect(() => {
        updateDescription()
    }, [description])

    function updateDescription() {
        if (!task) return
        api.patch(`/tasks/${task.uuid}`, {
            description
        }).then(res => {
            if (res.status < 300) {
                navigate("")
            } else {
                notifications.show({
                    title: "Failed to update description",
                    message: res.data.message,
                })
            }
        })
    }

    return (
        <div className={styles.wrapper}>
            <LoaderOverlay loading={true}>
                <textarea className={`${cx(classes.description)} ${styles.description}`} value={description} onChange={e => setDescription(e.target.value)}>
                    {task.description}
                </textarea>
            </LoaderOverlay>
        </div>
    )

}
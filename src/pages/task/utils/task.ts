import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {Task} from "@kanban/types2"


export function getTask(uuid: string) {
    const {environment} = useContext(EnvironmentContext)
    return environment?.tasks?.find((task) => task.uuid === uuid)
}

export function setTask(uuid: string, task: Task) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const index = environment?.tasks?.findIndex((task) => task.uuid === uuid)
    environment.tasks[index] = task
    setEnvironment(environment)
}
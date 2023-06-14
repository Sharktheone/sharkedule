import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {Tag} from "@kanban/types2"


export function getTag(uuid: string) {
    const {environment} = useContext(EnvironmentContext)
    return environment.tags.find((tag) => tag.uuid === uuid)
}

export function setTag(uuid: string, tag: Tag) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const index = environment.tags.findIndex((tag) => tag.uuid === uuid)
    environment.tags[index] = tag
    setEnvironment(environment)
}
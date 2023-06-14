import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {Board, Column} from "@kanban/types2"


export function getColumn(uuid: string) {
    const {environment} = useContext(EnvironmentContext)
    return environment.columns.find((column) => column.uuid === uuid)
}

export function setColumn(uuid: string, column: Column) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const index = environment.columns.findIndex((column) => column.uuid === uuid)
    environment.columns[index] = column

    setEnvironment(environment)

}
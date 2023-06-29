import {useContext} from "react"
import {EnvironmentContext} from "@kanban/environment"
import {Board} from "@kanban/types2"


export function getBoard(uuid: string) {
    const {environment} = useContext(EnvironmentContext)
    return environment?.boards.find((board) => board.uuid === uuid)
}

export function setBoard(uuid: string, board: Board) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const index = environment?.boards.findIndex((board) => board.uuid === uuid)
    environment.boards[index] = board
    setEnvironment(environment)
}
import {Container, createStyles, Text, Title} from "@mantine/core";
import {useStyles} from "./styles";
import Task from "./task/task"

export default function Column() {
    const {classes, cx} = useStyles();
    return (
        <Container className={cx(classes.column)}>
            <Title align="left" className={cx(classes.title)} order={3}>Column</Title>
            <Task/>
            <Task/>
            <Task/>
            <Task/>
            <Task/>
        </Container>
    )
}
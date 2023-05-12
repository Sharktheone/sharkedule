import {Container, createStyles, Text, Title} from "@mantine/core";
import {useStyles} from "./styles";

export default function Column() {
    const {classes, cx} = useStyles();
    return (
        <Container className={cx(classes.column)}>
            <Title align="left" className={cx(classes.title)} order={3}>Column</Title>

        </Container>
    )
}
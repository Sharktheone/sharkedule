import {Flex} from "@mantine/core"
import {useStyles} from "./styles"


export default function Task() {
    const {classes, cx} = useStyles();

    return (
        <Flex className={cx(classes.task)}>
            Task
        </Flex>
    )
}
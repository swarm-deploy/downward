import unittest

import downward


class LoadTest(unittest.TestCase):
    def test_load(self) -> None:
        info = downward.load(
            {
                downward.ENV_STACK_NAME: "billing",
                downward.ENV_SERVICE_ID: "service-id",
                downward.ENV_SERVICE_NAME: "billing_api",
                downward.ENV_TASK_ID: "task-id",
                downward.ENV_TASK_NAME: "billing_api.2.task-id",
                downward.ENV_TASK_SLOT: "2",
                downward.ENV_NODE_ID: "node-id",
                downward.ENV_NODE_NAME: "worker-1",
            }
        )

        self.assertEqual(info.stack.name, "billing")
        self.assertEqual(info.service.id, "service-id")
        self.assertEqual(info.service.name, "billing_api")
        self.assertEqual(info.task.id, "task-id")
        self.assertEqual(info.task.name, "billing_api.2.task-id")
        self.assertEqual(info.task.slot, 2)
        self.assertEqual(info.node.id, "node-id")
        self.assertEqual(info.node.name, "worker-1")

    def test_load_without_slot(self) -> None:
        info = downward.load({})

        self.assertIsNone(info.task.slot)

    def test_load_with_invalid_slot(self) -> None:
        with self.assertRaisesRegex(ValueError, "parse SWARM_TASK_SLOT"):
            downward.load({downward.ENV_TASK_SLOT: "invalid"})

    def test_to_dict(self) -> None:
        info = downward.Info(
            stack=downward.Stack(name="billing"),
            service=downward.Service(id="service-id", name="billing_api"),
            task=downward.Task(id="task-id", name="billing_api.2.task-id", slot=2),
            node=downward.Node(id="node-id", name="worker-1"),
        )

        self.assertEqual(
            info.to_dict(),
            {
                downward.ENV_STACK_NAME: "billing",
                downward.ENV_SERVICE_ID: "service-id",
                downward.ENV_SERVICE_NAME: "billing_api",
                downward.ENV_TASK_ID: "task-id",
                downward.ENV_TASK_NAME: "billing_api.2.task-id",
                downward.ENV_TASK_SLOT: 2,
                downward.ENV_NODE_ID: "node-id",
                downward.ENV_NODE_NAME: "worker-1",
            },
        )


class ServiceTest(unittest.TestCase):
    def test_short_name(self) -> None:
        service = downward.Service(id="service-id", name="billing_api")

        self.assertEqual(service.short_name("billing"), "api")

    def test_short_name_without_stack(self) -> None:
        service = downward.Service(id="service-id", name="billing_api")

        self.assertEqual(service.short_name(""), "billing_api")


if __name__ == "__main__":
    unittest.main()

# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: qtaskd.proto
# plugin: python-betterproto
from dataclasses import dataclass
from datetime import datetime
from typing import Dict

import betterproto
import grpclib
from betterproto.grpc.grpclib_server import ServiceBase


@dataclass(eq=False, repr=False)
class Request(betterproto.Message):
    message: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class Reply(betterproto.Message):
    message: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class RunTaskRequest(betterproto.Message):
    id: str = betterproto.string_field(1)
    status: str = betterproto.string_field(2)
    created_at: datetime = betterproto.message_field(3)
    started_at: datetime = betterproto.message_field(4)
    paused_at: datetime = betterproto.message_field(5)
    terminated_at: datetime = betterproto.message_field(6)
    name: str = betterproto.string_field(7)
    description: str = betterproto.string_field(8)
    working_dir: str = betterproto.string_field(9)
    command_line: str = betterproto.string_field(10)
    output_file_path: str = betterproto.string_field(11)


@dataclass(eq=False, repr=False)
class GetTaskRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class GetTaskReply(betterproto.Message):
    id: str = betterproto.string_field(1)
    status: str = betterproto.string_field(2)


class QTaskDaemonStub(betterproto.ServiceStub):
    async def echo(self, *, message: str = "") -> "Reply":

        request = Request()
        request.message = message

        return await self._unary_unary("/qtaskd.QTaskDaemon/Echo", request, Reply)

    async def run_task(
            self,
            *,
            id: str = "",
            status: str = "",
            created_at: datetime = None,
            started_at: datetime = None,
            paused_at: datetime = None,
            terminated_at: datetime = None,
            name: str = "",
            description: str = "",
            working_dir: str = "",
            command_line: str = "",
            output_file_path: str = "",
    ) -> "Reply":

        request = RunTaskRequest()
        request.id = id
        request.status = status
        if created_at is not None:
            request.created_at = created_at
        if started_at is not None:
            request.started_at = started_at
        if paused_at is not None:
            request.paused_at = paused_at
        if terminated_at is not None:
            request.terminated_at = terminated_at
        request.name = name
        request.description = description
        request.working_dir = working_dir
        request.command_line = command_line
        request.output_file_path = output_file_path

        return await self._unary_unary("/qtaskd.QTaskDaemon/RunTask", request, Reply)

    async def get_task(self) -> "GetTaskReply":

        request = GetTaskRequest()

        return await self._unary_unary(
            "/qtaskd.QTaskDaemon/GetTask", request, GetTaskReply
        )


class QTaskDaemonBase(ServiceBase):
    async def echo(self, message: str) -> "Reply":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def run_task(
            self,
            id: str,
            status: str,
            created_at: datetime,
            started_at: datetime,
            paused_at: datetime,
            terminated_at: datetime,
            name: str,
            description: str,
            working_dir: str,
            command_line: str,
            output_file_path: str,
    ) -> "Reply":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get_task(self) -> "GetTaskReply":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_echo(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "message": request.message,
        }

        response = await self.echo(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_run_task(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "id": request.id,
            "status": request.status,
            "created_at": request.created_at,
            "started_at": request.started_at,
            "paused_at": request.paused_at,
            "terminated_at": request.terminated_at,
            "name": request.name,
            "description": request.description,
            "working_dir": request.working_dir,
            "command_line": request.command_line,
            "output_file_path": request.output_file_path,
        }

        response = await self.run_task(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_get_task(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.get_task(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/qtaskd.QTaskDaemon/Echo": grpclib.const.Handler(
                self.__rpc_echo,
                grpclib.const.Cardinality.UNARY_UNARY,
                Request,
                Reply,
            ),
            "/qtaskd.QTaskDaemon/RunTask": grpclib.const.Handler(
                self.__rpc_run_task,
                grpclib.const.Cardinality.UNARY_UNARY,
                RunTaskRequest,
                Reply,
            ),
            "/qtaskd.QTaskDaemon/GetTask": grpclib.const.Handler(
                self.__rpc_get_task,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetTaskRequest,
                GetTaskReply,
            ),
        }

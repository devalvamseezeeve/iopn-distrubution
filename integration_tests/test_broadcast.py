from pathlib import Path

import pytest

from .network import setup_custom_iopn
from .utils import submit_any_proposal

pytestmark = pytest.mark.gov


@pytest.fixture(scope="module")
def custom_iopn(tmp_path_factory):
    path = tmp_path_factory.mktemp("iopn")
    yield from setup_custom_iopn(
        path, 26400, Path(__file__).parent / "configs/broadcast.jsonnet"
    )


def test_submit_any_proposal(custom_iopn, tmp_path):
    submit_any_proposal(custom_iopn, tmp_path)

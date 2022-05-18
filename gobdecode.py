import sys, os
from datetime import datetime
from typing import BinaryIO, Generator, Any
from uuid import uuid4
import pygob

from nfcgate_proto.c2c_pb2 import NFCData
from nfcgate_proto.c2s_pb2 import ServerData


def hex_dump(a: bytes) -> str:
    # return " ".join([f"0x{str(hex(c))[2:].zfill(2)}" for c in a])
    return " ".join([f"{str(hex(c))[2:].zfill(2)}" for c in a])


def format_data(f: BinaryIO) -> None:
    encoded_data = f.read()
    f.close()

    output_file = f"{f.name.split('.')[0]}-decoded.log"
    
    if os.path.exists(output_file):
        print(f"Output file {output_file} already exists!")
        output_file = "".join((output_file.split('.')[0] + str(uuid4()), ".log"))
        

    print(f"Writing to {output_file} ...")

    gob = pygob.Loader()
    decoded_gen: Generator[Any, None, None] = gob.load_all(encoded_data)

    for decoded_data in decoded_gen:
        data = decoded_data.Data

        if len(data) == 0:
            continue

        server_data: ServerData = ServerData()
        server_data.ParseFromString(data)

        nfc_data: NFCData = NFCData()
        nfc_data.ParseFromString(server_data.data)

        letter: str = "Tag" if nfc_data.data_source == NFCData.CARD else "Rdr"
        initial: str = "(initial) " if nfc_data.data_type == NFCData.INITIAL else ""
        string_to_return: str = f"{datetime.fromtimestamp(nfc_data.timestamp/1000)} | {letter} | {hex_dump(nfc_data.data)} {'| ' + initial if initial else ''}\n"
            
        with open(output_file, "a") as f:
            f.write(string_to_return)
    
    print(f"Written logs to {output_file}!")


def main() -> None:
    if len(sys.argv) < 2:
        print(f"Usage: python3 {sys.argv[0]} <filename.gob>")
        sys.exit(-1)

    gobfile: str = sys.argv[1]

    f: BinaryIO = open(gobfile, "rb")
    format_data(f)

if __name__ == "__main__":
    main()
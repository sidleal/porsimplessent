# PorSimplesSent
## A Portuguese corpus of aligned sentences pairs to investigate sentence readability assessment

## NILC
This corpus was created during my master degree at ICMC-USP, and made possible thanks to the Interinstitutional Center for Computational Linguistics - NILC (Núcleo Interestitucional de Linguística Computacional), represented by my advisor Dra. Sandra Maria Aluísio and the linguistics specialist Dra. Magali Sanches Duran.

[http://www.nilc.icmc.usp.br/nilc/index.php](http://www.nilc.icmc.usp.br/nilc/index.php)


## License
[https://creativecommons.org/licenses/by/4.0/](CC BY 4.0)

## TSV format
All files are in Tab Separated Values (TSV) format, it means that fields are separated by tab (Also knows as ```char(9)``` or ```\t```), and newline (```char(10)``` or ```\n```) for the rows.

## porsimples

In this folder you'll find the source corpus used to extract the sentence pairs, already exportaded in TSV format:

### porsimples_sentences.tsv
* production_id: Each triplet of texts (original, natural, strong) has an unique id, called production_id.
* level: ORI (1 - Original), level NAT (2 - Natural) or STR (3 - Strong)
* text_id: Unique id for each text.
* sentence_id: Unique id for each sentence.
* paragraph: Sequential id for the paragraph in text.
* sentence_text: The raw text from the sentence.

### porsimples_aligns.tsv
* production_id: See porsimples_sentences.tsv.
* level: Simplification level ORI->NAT or NAT->STR.
* text_id_from: Text id from source side of simplification.
* sentence_id_from: sentence id from source side of simplification.
* text_id_to: Text id for target side of simplification.
* sentence_id_to: Sentence id for target side of simplification.

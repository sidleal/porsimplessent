# PorSimplesSent
## A Portuguese corpus of aligned sentences pairs to investigate sentence readability assessment

## NILC
This corpus was created during my master's degree at ICMC-USP, and made possible thanks to the Interinstitutional Center for Computational Linguistics - NILC (Núcleo Interestitucional de Linguística Computacional), represented by my advisor Dra. Sandra Maria Aluísio and the linguistics specialist Dra. Magali Sanches Duran.

[http://www.nilc.icmc.usp.br/nilc/index.php](http://www.nilc.icmc.usp.br/nilc/index.php)


## License
[CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)

## TSV format
All files are in Tab Separated Values (TSV) format, it means that fields are separated by tab (Also knows as ```char(9)``` or ```\t```), and newline (```char(10)``` or ```\n```) for the rows.

## porsimples

In this folder you'll find the source corpus used to extract the sentence pairs, already exportaded in TSV format:

### porsimples_sentences.tsv
* __production_id__: Each triplet of texts (original, natural, strong) has an unique id, called production_id.
* __level__: ORI (1 - Original), level NAT (2 - Natural) or STR (3 - Strong)
* __text_id__: Unique id for each text.
* __sentence_id__: Unique id for each sentence.
* __paragraph__: Sequential id for the paragraph in text.
* __sentence_text__: The raw text from the sentence.

### porsimples_aligns.tsv
* __production_id__: See porsimples_sentences.tsv.
* __level__: Simplification level ORI->NAT or NAT->STR.
* __text_id_from__: Text id from source side of simplification.
* __sentence_id_from__: sentence id from source side of simplification.
* __text_id_to__: Text id for target side of simplification.
* __sentence_id_to__: Sentence id for target side of simplification.
